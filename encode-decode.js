const polyline = require('@mapbox/polyline');

function encode(coords) {
    return polyline.encode(coords);
}

function decode(encoded) {
    return polyline.decode(encoded);
}

let encoded = encode([[37.7749, -122.4194], [34.0522, -118.2437]]);
console.log("Encoded: ", encoded);

let decoded = decode(encoded);
console.log("Decoded: ", decoded);
