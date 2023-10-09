import polyline

def encode(coords):
    return polyline.encode(coords, 5)

def decode(encoded):
    return polyline.decode(encoded, 5)

encoded = encode([(37.7749, -122.4194), (34.0522, -118.2437)]) #Refer to encoded/decoded polylines from https://github.com/mapup/Encode-Decode-Google-Polyline/tree/main/sample-polylines
print("Encoded: ", encoded)

decoded = decode(encoded) #Refer to encoded/decoded polylines from https://github.com/mapup/Encode-Decode-Google-Polyline/tree/main/sample-polylines
print("Decoded: ", decoded)