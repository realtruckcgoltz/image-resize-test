from PIL import Image
import multiprocessing

def worker(num):
    size = 128, 128
    im = Image.open("../images/earth.jpg")
    im.thumbnail(size)
    im.save("thumbs/earth-" + str(num) + ".jpg", "JPEG")

def main():
    jobs = []
    for i in range(50):
        p = multiprocessing.Process(target=worker, args=(i,))
        jobs.append(p)
        p.start()

if __name__ == "__main__":
    main()