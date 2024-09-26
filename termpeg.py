import sys
import math
import numpy as np
import os
from PIL import Image




size = 1920, 1080

infile = sys.argv[1]
outfile = os.path.splitext(infile)[0] + ".thumbnail"
if infile != outfile:
    try:
        im = Image.open(infile)
        im.thumbnail(size, Image.Resampling.LANCZOS)
        im.save(outfile, "JPEG")
    except IOError:
        print(f"cannot create thumbnail for {infile}")

img = Image.open(outfile)
arr = np.array(img)
shape = arr.shape
print(shape)

def pixel_to_byte(px):
    # Extracting the RGB values
    r, g, b = px[0], px[1], px[2]
    
    # Conversion logic
    r_byte = px[0] & 0b11100000
    g_byte = px[1] >> 6
    b_byte = px[2] >> 5

    
    # if r_byte >= 8:

    # Combine into a single byte
    byte_val = r_byte + (g_byte << 3) + b_byte
    
    return byte_val

temp = []
for i in range(shape[0]):
   for j in range(shape[1]):
        temp.append(pixel_to_byte(arr[i][j]))


output_file = "akane.vido"
with open(output_file, "wb") as f:
    newFileByteArray = bytearray(temp)
    f.write(newFileByteArray)

print(f"Binary data written to {output_file}")


