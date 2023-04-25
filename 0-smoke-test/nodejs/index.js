const net = require('node:net');

const server = net.createServer(c => {
    console.log('client connected');

    let input = [];

    c.on('data', (s) => {
        console.log('received data');
        console.log(input.length);
        input.push(s);
    })

    c.on('end', () => {
        console.log('end');
        c.write(Buffer.concat(input));
    });

    c.on('close', () => {
        console.log('close');
    });
});

server.listen(3333, () => {
    console.log('server is listening');
});
