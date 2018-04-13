from PIL import Image

size = 128, 128

im = Image.open("../images/earth.jpg")
im.thumbnail(size)
im.save("thumbs/earth.jpg", "JPEG")