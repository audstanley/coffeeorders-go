# Coffeeorders-go

is an API for [Front-End Web Development: The Big Nerd Ranch Guide](https://www.bignerdranch.com/books/front-end-web-development-the-big-nerd-ranch-guide/)
The book is missing some endpoints, and this GoLang Project recreates those endpoints and is compiled for windows, mac, and linux.


## Best way to run the application:

Check [Releases](https://github.com/audstanley/coffeeorders-go/releases). Download the bin.zip, extract, and run the binary for your system.

If youre on Linux/Mac, you'll likely need to do these things after you uncompress the zip:


On Mac:

```bash
# navigate into the unziped bin folder
chmod +x coffeeorders-amd64-macos
./coffeeorders-amd64-macos
# then the program should run
# this assumes you are not on a new mac with the arm64 processors
# if it does work on the new macs, please open an issue on github and let me know
# if it does or does not work.  I would be interested in knowing
```

On Linux
```bash
# navigate into the unziped bin folder
chmod +x coffeeorders-amd64-linux
./coffeeorders-amd64-linux
# then the program should run
```

On Windows:
```bash
# depending on if you windows is x86, or x64 (most likely x64)
# navigate into the unziped bin folder
coffeeorders-amd64.exe
```

### Once you have the binaries running, you will want this extension:

[REST](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)

Once you install the rest extension, you can copy the .http files that are in this repositorie's rest folder,
don't forget to copy the rest/.env file, since that is where are the http calls are assigned to.

In the rest/.env
```bash
baseUrl=https://co.audstanley.com
# can be changed to
baseUrl=http://127.0.0.1:3001
# and now all the http calls will be made to your own computer.
```

## TLDR video;

[![tldr](https://i9.ytimg.com/vi/umPXiItPzbQ/mq1.jpg?sqp=CJzVz4oG&rs=AOn4CLAng2mGYz3XOKWjsiJ3AWni9XnCwQ)](https://youtu.be/umPXiItPzbQ)




If there is something wrong with the API, please feel free to contibute code or open an issue.

The Current builds are for amd64 processors, and also windows x86.  If you need a different archetecture, let me know. I can add to a future release.


You are welcome to:

<a href='https://ko-fi.com/A687KA8' target='_blank'><img height='36' style='border:0px;height:36px;' src='https://az743702.vo.msecnd.net/cdn/kofi4.png?v=f' border='0' alt='Buy Me a Coffee at ko-fi.com' /></a>

Or stop by my blog [audstanley.com](https://www.audstanley.com)

### If you want to run the application as a Docker container - or help to develop - a docker container can be useful:

```bash
# you can just run the run-dockerfile.sh if you are on linux and have docker installed.
# if you are not on linux, you can still just run lines 2, and 3 from within that script,
# and from within your project folder like so:
docker build -t coffee-orders .
docker run -dp 3001:3001 coffee-orders
```

profit.