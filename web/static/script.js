let map;
let polyline;
let currentIndex = 0;
let isPlaying = true;
let animationSpeed = 1.0;
let animationFrameId = null;
let currentDirection = 0;
let lastDirectionUpdate = 0;
let coordinates = [];
let timestamps = [];

// Initialize the map
function initMap() {
    map = L.map('map').setView([57.1631414, 65.5686189], 15);
    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '© OpenStreetMap contributors'
    }).addTo(map);

    // Create a moving arrow marker
    const arrowIcon = L.divIcon({
        className: 'arrow-marker',
        html: '<div style="width: 20px; height: 10px; background: #2ecc71; border: 1px solid #27ae60; clip-path: polygon(0% 0%, 100% 50%, 0% 100%); transform: rotate(0deg);"></div>',
        iconSize: [20, 10],
        iconAnchor: [0, 5]
    });

    window.movingArrow = L.marker([0, 0], {
        icon: arrowIcon,
        zIndexOffset: 1000
    }).addTo(map);
}

// Load geo data from the server
async function loadGeoData() {
    try {
        const response = await fetch('/api/geo-data');
        const data = await response.json();
        
        // Sort logs by timestamp
        const sortedLogs = data.geo_position_logs.sort((a, b) => a.timestamp - b.timestamp);
        coordinates = sortedLogs.map(log => [log.latitude, log.longitude]);
        timestamps = sortedLogs.map(log => log.timestamp);
        
        // Create the track line
        polyline = L.polyline(coordinates, {
            color: '#ff0000',
            weight: 5,
            opacity: 0.5
        }).addTo(map);
        
        // Add points
        coordinates.forEach((coord, index) => {
            const point = L.circleMarker(coord, {
                radius: 4,
                fillColor: '#ff0000',
                color: '#000',
                weight: 1,
                opacity: 0,
                fillOpacity: 0.75
            }).addTo(map);
            
            const date = new Date(timestamps[index]);
            point.bindPopup(`Point ${index + 1}<br>Time: ${date.toLocaleTimeString()}`);
        });
        
        // Start animation
        animateTrack();
    } catch (error) {
        console.error('Error loading geo data:', error);
    }
}

// Animation functions
function transformSpeed(sliderValue) {
    const sign = Math.sign(sliderValue);
    const absValue = Math.abs(sliderValue);
    return sign * Math.pow(absValue / 25, 1) * 25;
}

function getCarDirection(start, end) {
    function toRadians(degrees) {
        return degrees * Math.PI / 180;
    }
    
    function toDegrees(radians) {
        return radians * 180 / Math.PI;
    }
    
    const startLat = toRadians(start.lat);
    const startLng = toRadians(start.lng);
    const destLat = toRadians(end.lat);
    const destLng = toRadians(end.lng);
    
    const y = Math.sin(destLng - startLng) * Math.cos(destLat);
    const x = Math.cos(startLat) * Math.sin(destLat) -
              Math.sin(startLat) * Math.cos(destLat) * Math.cos(destLng - startLng);
    let brng = Math.atan2(y, x);
    brng = toDegrees(brng);
    return (brng + 360) % 360;
}

function haversineDistance(coord1, coord2) {
    const R = 6371e3;
    const φ1 = coord1[0] * Math.PI/180;
    const φ2 = coord2[0] * Math.PI/180;
    const Δφ = (coord2[0] - coord1[0]) * Math.PI/180;
    const Δλ = (coord2[1] - coord1[1]) * Math.PI/180;
    
    const a = Math.sin(Δφ/2) * Math.sin(Δφ/2) +
              Math.cos(φ1) * Math.cos(φ2) *
              Math.sin(Δλ/2) * Math.sin(Δλ/2);
    const c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1-a));
    
    return R * c;
}

function animateTrack() {
    let startTime = null;
    let currentTimestamp = timestamps[0];
    let lastIndex = 0;
    
    function findNearestIndex(targetTimestamp) {
        let left = 0;
        let right = timestamps.length - 1;
        
        while (left <= right) {
            const mid = Math.floor((left + right) / 2);
            if (timestamps[mid] === targetTimestamp) {
                return mid;
            } else if (timestamps[mid] < targetTimestamp) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        
        if (left >= timestamps.length) return timestamps.length - 1;
        if (right < 0) return 0;
        
        return Math.abs(timestamps[left] - targetTimestamp) < 
               Math.abs(timestamps[right] - targetTimestamp) ? left : right;
    }
    
    function animate(currentTime) {
        if (!startTime) startTime = currentTime;
        const elapsed = currentTime - startTime;
        
        const timeIncrement = (elapsed * animationSpeed) / 1000;
        const newTimestamp = currentTimestamp + timeIncrement;
        
        if (newTimestamp > timestamps[timestamps.length - 1]) {
            const lastIndex = timestamps.length - 1;
            window.movingArrow.setLatLng(coordinates[lastIndex]);
            map.setView(coordinates[lastIndex]);
            currentTimestamp = timestamps[lastIndex];
        } else if (newTimestamp < timestamps[0]) {
            window.movingArrow.setLatLng(coordinates[0]);
            map.setView(coordinates[0]);
            currentTimestamp = timestamps[0];
        } else {
            const newIndex = findNearestIndex(newTimestamp);
            const newPosition = coordinates[newIndex];
            window.movingArrow.setLatLng(newPosition);
            map.setView(newPosition);
            
            if (newIndex !== lastIndex) {
                const nextIndex = Math.min(newIndex + 1, coordinates.length - 1);
                const rotation = 90 + getCarDirection(
                    { lat: coordinates[newIndex][0], lng: coordinates[newIndex][1] },
                    { lat: coordinates[nextIndex][0], lng: coordinates[nextIndex][1] }
                );
                
                const hd = haversineDistance(coordinates[newIndex], coordinates[nextIndex]);
                if (hd > 4) {
                    currentDirection = rotation;
                    lastDirectionUpdate = nextIndex;
                }
                
                const arrowElement = window.movingArrow.getElement();
                if (arrowElement) {
                    arrowElement.querySelector('div').style.transform = `rotate(${currentDirection}deg)`;
                }
                lastIndex = newIndex;
            }
            
            currentTimestamp = newTimestamp;
        }
        
        if (isPlaying) {
            animationFrameId = requestAnimationFrame(animate);
        }
    }
    
    animationFrameId = requestAnimationFrame(animate);
}

// Event listeners
document.addEventListener('DOMContentLoaded', () => {
    initMap();
    loadGeoData();
    
    const speedSlider = document.getElementById('speedSlider');
    const speedValue = document.getElementById('speedValue');
    const pauseButton = document.getElementById('pauseButton');
    
    speedSlider.addEventListener('input', (e) => {
        animationSpeed = transformSpeed(parseFloat(e.target.value));
        speedValue.textContent = `${e.target.value}x`;
    });
    
    pauseButton.addEventListener('click', () => {
        isPlaying = !isPlaying;
        pauseButton.textContent = isPlaying ? 'Pause' : 'Play';
        if (isPlaying) {
            animateTrack();
        } else {
            cancelAnimationFrame(animationFrameId);
        }
    });
    
    document.getElementById('resetButton').addEventListener('click', () => {
        currentIndex = 0;
        currentTimestamp = timestamps[0];
        window.movingArrow.setLatLng(coordinates[0]);
        map.setView(coordinates[0]);
    });
}); 