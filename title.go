package main

var (
	Css = `<style type='text/css'>
	a:hover {
	color:black;
	}
	a:visited {
	color:black;
	}
	a:link {
	color:black;
	}
	a {
	color:black;
	}
	
	/* Border styles */
	table thead, table tr {
	border-top-width: 1px;
	border-top-style: solid;
	border-top-color: #cccccc;
	}
	table thead {
		background: #c4d6f9
	}
	table {
	border-bottom-width: 1px;
	border-bottom-style: solid;
	border-bottom-color: #ffffff;
	}
	
	/* Padding and font style */
	table td, table th {
	padding: 5px 10px;
	font-size: 12px;
	font-family: Verdana;
	color: #444444;
	}
	
	/* Alternating background colors */
	table tbody tr:nth-child(even) {
	background: #f0f0f0
	}
	table tbody tr:nth-child(odd) {
	background: #FFF
	}
		table tbody tr:hover {
		background: #dddddd
	}
	table tbody tr:hover td {
	background:none;
	}</style>`

	Meta = `<meta name="viewport" content="width=device-width, initial-scale=0.9, user-scalable=no, minimum-scale=0.9, maximum-scale=0.9" />  
<meta name="apple-mobile-web-app-capable" content="yes" />  
<meta name="format-detection" content="telephone=no" />`

	Title = Css + Meta + `<h1><a href='/'>webserv @aoaolion</a></h1>
		
		<a href='/upload'>upload</a> &nbsp;&nbsp;
		<a href='/download'>download</a> &nbsp;&nbsp;
		<a href='/close'>close</a> &nbsp;&nbsp;
		<a href='https://github.com/aoaolion/webserv' target='_blank'>https://github.com/aoaolion/webserv</a><br><br>`
)
