const sharp = require('sharp')

sharp('../images/earth.jpg')
    .resize(128, 128)
    .max()
    .toFile('thumbs/earth.jpg', (err) => {})
