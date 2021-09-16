# Coffeeorders-go

is an API for [Front-End Web Development: The Big Nerd Ranch Guide](https://www.bignerdranch.com/books/front-end-web-development-the-big-nerd-ranch-guide/)
The book is missing some endpoints, and this GoLang Project recreates those endpoints and is compiled for windows, mac, and linux.

Check [Releases](https://github.com/audstanley/coffeeorders-go/releases). Download the bin.zip, extract, and run the binary for your system.

If there is something wrong with the API, please feel free to contibute code or open an issue.

The Current builds are for amd64 processors, and also windows x86.  If you need a different archetecture, let me know. I can add to a future release.


You are welcome to:

<a href='https://ko-fi.com/A687KA8' target='_blank'><img height='36' style='border:0px;height:36px;' src='https://az743702.vo.msecnd.net/cdn/kofi4.png?v=f' border='0' alt='Buy Me a Coffee at ko-fi.com' /></a>

Or stop by my blog [audstanley.com](https://www.audstanley.com)

### If you want to run the application as a Docker container:

```bash
# you can just run the run-dockerfile.sh if you are on linux and have docker installed.
# if you are not on linux, you can still just run lines 2, and 3 from within that script,
# and from within your project folder like so:
docker build -t coffee-orders .
docker run -dp 3000:3000 coffee-orders
```

profit.