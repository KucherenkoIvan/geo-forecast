package com.example.geoforecast

import android.Manifest
import android.os.Bundle
import android.os.Looper
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge
import androidx.annotation.RequiresPermission
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.tooling.preview.Preview
import com.example.geoforecast.ui.theme.GeoForecastTheme
import com.google.android.gms.location.FusedLocationProviderClient
import com.google.android.gms.location.LocationCallback
import com.google.android.gms.location.LocationRequest
import com.google.android.gms.location.LocationResult
import com.google.android.gms.location.LocationServices
import java.net.HttpURLConnection
import java.net.URL


class MainActivity : ComponentActivity() {
    private var isOnline = false;
    private lateinit var fusedLocationClient: FusedLocationProviderClient
    private lateinit var locationCallback: LocationCallback

    fun sendLocationUpdate(lat:Number, lng:Number) {
        val host = "138.124.102.191"
        val port = "9038"
        val mURL = URL("http://${host}:${port}/api/position_log")

        try {
            with(mURL.openConnection() as HttpURLConnection) {
                // optional default is GET
                requestMethod = "POST"
                addRequestProperty("Authorization", "Bearer testandroidapp")
                setRequestProperty("Content-Type", "application/json");
                setDoOutput(true);
                val jsonInputString = "{ \"lat\": $lat, \"lng\": $lng }"

                getOutputStream().use { os ->
                    val input = jsonInputString.toByteArray(charset("utf-8"))
                    os.write(input, 0, input.size)
                }

                println("URL : $url")
                println("Response Code : $responseCode")
                isOnline = true;
            }
        } catch (e:Exception) {
            this.isOnline = true
            println(e)
        }
    }

    @RequiresPermission(anyOf = [Manifest.permission.ACCESS_FINE_LOCATION, Manifest.permission.ACCESS_COARSE_LOCATION])
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        this.fusedLocationClient = LocationServices.getFusedLocationProviderClient(this)

        locationCallback = object : LocationCallback() {
            override fun onLocationResult(p0: LocationResult) {
                val lat = p0.lastLocation?.latitude
                val lng = p0.lastLocation?.longitude
                val utime = p0.lastLocation?.time
                val mt = "{\n    \"lat\": ${lat ?: "null"},\n    \"lng\": ${lng ?: "null"},\n    \"utime\": ${utime ?: "null"}\n}"

                if (lat != null && lng !== null) {
                    sendLocationUpdate(lat, lng)
                }

                setContent {
                    GeoForecastTheme {
                        Scaffold(modifier = Modifier.fillMaxSize()) { innerPadding ->
                            Greeting(
                                isOnline,
                                name = mt,
                                modifier = Modifier.padding(innerPadding)
                            )
                        }
                    }
                }
            }
        }
        startLocationUpdates()

        enableEdgeToEdge()
        setContent {
            GeoForecastTheme {
                Scaffold(modifier = Modifier.fillMaxSize()) { innerPadding ->
                    Greeting(
                        isOnline,
                        name = "null",
                        modifier = Modifier.padding(innerPadding)
                    )
                }
            }
        }
    }
    @RequiresPermission(anyOf = [Manifest.permission.ACCESS_FINE_LOCATION, Manifest.permission.ACCESS_COARSE_LOCATION])
    private fun startLocationUpdates() {
        this.fusedLocationClient.requestLocationUpdates(
            LocationRequest.Builder(100)
                .setPriority(LocationRequest.PRIORITY_HIGH_ACCURACY)
                .setIntervalMillis(0L)
                .setMinUpdateIntervalMillis(0L)
                .build(),
            locationCallback,
            Looper.getMainLooper())
    }
}


@Composable
fun Greeting(isOnline: Boolean, name: String, modifier: Modifier = Modifier) {
    var text = "OFFLINE"
    if (isOnline) {
        text = "ONLINE"
    }

    Text(
        text = "Status: $text;\nLast location: $name",
        modifier = modifier
    )
}

@Preview(showBackground = true)
@Composable
fun GreetingPreview() {
    GeoForecastTheme {
        Greeting(false, "null")
    }
}
