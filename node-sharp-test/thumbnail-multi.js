const sharp = require('sharp')

for (let i = 0; i < 50; i++) {
    sharp('../images/earth.jpg')
        .resize(128, 128)
        .max()
        .toFile('thumbs/' + i + '.jpg', (err) => {})
}