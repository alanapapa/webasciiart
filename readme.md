<h2>Hello, I'm Berik <i>Alanapapa</i> Bazarov!</h2>
<img align='right' src="gopher-rainbow.gif" width="230">
<p><em>Freelancer and student at <a href="https://alem.school/" target="_blank">Alem Coding School</a><img src="https://media.giphy.com/media/WUlplcMpOCEmTGBtBW/giphy.gif" width="30"> 
</em></p>

[![Github: alanapapa](https://img.icons8.com/material-outlined/48/000000/github.png?style=flat-square&logo=Instagram&logoColor=white&link=https://www.github.com/alanapapa/)](https://www.github.com/alanapapa/)
[![Instagram: berikbazarov](https://img.icons8.com/fluent/48/000000/instagram-new.png?style=flat-square&logo=Instagram&logoColor=white&link=https://www.instagram.com/berikbazarov/)](https://www.instagram.com/berikbazarov/)
[![Facebook: berikbazarov](https://img.icons8.com/color/48/000000/facebook.png?style=flat-square&logo=Instagram&logoColor=white&link=https://www.facebook.com/bazarovberik/)](https://www.facebook.com/bazarovberik/)

![GitHub followers](https://img.shields.io/github/followers/alanapapa?label=Follow&style=social)


## ***ascii-art-web***
### This is my solution to the *ascii-art-web* educational task on [01 Edu System](https://www.01-edu.org/)

```console
student@ubuntu:~/ascii-art-web$ go build
student@ubuntu:~/ascii-art-web$ ./webasciiart
Starting server for testing HTTP POST...
 http://127.0.0.1:1995/ 

```
![image](screen.png)

<br>

### Objectives

Ascii-art-web consists in creating and running a server, in which it will be possible to use a web **GUI** (graphical user interface) version of your last project, [ascii-art](https://public.01-edu.org/subjects/ascii-art/).

Your web-page should provide usage of different [banners](https://github.com/01-edu/public/tree/master/subjects/ascii-art).

Implement following HTTP endpoints:

1. GET `/`: Sends HTML response - the main page.\
 1.1. GET Tip: [go templates](https://golang.org/pkg/html/template/) to receive and display data from the server

2. POST `/ascii-art`: that sends data to Go server (text and a banner)\
 2.1. POST Tip: use form and other types of [tags](https://developer.mozilla.org/en-US/docs/Web/HTML/Element) to make the post request.\

Main page must have:

- text input
- radio buttons, select object or anything else to choose between banners
- button, which sends a POST request to '/ascii-art' and outputs the result on page.

### HTTP status code

Your endpoints must return appropriate HTTP status codes.

- OK (200), if everything went without errors
- Not Found, if anything is not found, e.g: template, banner etc.
- Bad Request, for incorrect requests
- Internal Server Error, for unhandled errors

## Markdown

In root project directory create `README.MD` file with the following sections and contents:

- Description
- Authors
- Usage: how to run
- Implementation details: algorithm

### Allowed packages

- Only the [standard go](https://golang.org/pkg/) packages are allowed

### Instructions

- HTTP server must be written in _Go_.
- HTML templates must be in project root directory _templates_.
- The code must respect the [good practices](https://public.01-edu.org/subjects/good-practices/).

### Usage

- [Here's an example](http://patorjk.com/software/taag/#p=display&f=Graffiti&t=Type%20Something%20).
