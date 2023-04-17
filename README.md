# Disable Touchpad

<p>
I had some irritated actions when typing message to any one or when4
just search in internet. I tried some libs for Linux to avoid this problem,
but it was not working correctly. So I decided write dispatcher to 
control when enable and disable touchpad.
</p>

<h2> Installation </h2>
<p>to install by yourself you need do:</p>
<ul>
<li>Download from github <b>disable-touchpad</b></li>
<li>Download and install Go version 1.20</li>
<li>Enter root folder of project</li>
<li>Run command <code>go mod tidy</code> to download dependencies</li>
<li>Run command <code>go build cmd/disable-touchpad/disable-touchpad.go</code> to build binary</li>
<li>Finnaly you can use by just enter in command line <code>./disable-touchpad</code></li>
</ul>