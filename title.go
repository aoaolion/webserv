package main

var (
	Css = `<style type='text/css'>
		body, html, div {
		margin: 0px;
		padding: 0px;
	}
	body {
		background-color: #eeeeee
	}
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
		text-decoration:none;
	}
	.clearfix:before, .clearfix:after {
		content: "";
		display: table;
	} 
	.clearfix:after {
		clear: both;
	}
	table thead, table tr {
		border-top-width: 1px;
		border-top-style: solid;
		border-top-color: #cccccc;
	}
	table thead {
		#background: #a5b7db
	}
	table {
		border-bottom-width: 1px;
		border-bottom-style: solid;
		border-bottom-color: #ffffff;
	}
	table td, table th {
		padding: 5px 10px;
		font-size: 12px;
		font-family: Verdana;
		color: #444444;
	}
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
	}
	ul{
		list-style-type:none;
		margin:0px 0px;/*margin:100px auto无效,不能使ul左右居中*/
		text-align:center;
		font-size:14px;
		padding-left: 0px;
	}
	li{
		float:left;/*改动的地方*/
		width:80px;
		padding:10px;
		background-color:#3c3d45;
	}
	ul a:link,ul a:visited,ul a:hover,ul a:active{
		color:#fff;
		text-decoration:none;
	}
	.btn {
		border: 1px solid #000;
		padding: 2px 3px;
		border-radius:5px;
		width: 50px;
		float: left;
	}
</style>`

	MetaDevice = `<meta name="viewport" content="width=device-width, initial-scale=0.9, user-scalable=no, minimum-scale=0.9, maximum-scale=0.9" />  
<meta name="apple-mobile-web-app-capable" content="yes" />  
<meta name="format-detection" content="telephone=no" />`

	MetaRedirect = `<meta http-equiv=refresh content='1;url=/'>`

	Head         = "<head>" + MetaDevice + Css + "</head>"
	HeadRedirect = "<head>" + MetaDevice + MetaRedirect + Css + "</head>"

	Nav = `<header style="width: 100%; height: 42px; background: #3c3d45; position: relative;">
		<div style="position: absolute;left:0;top:0;">
			<p style='font-size: 30px;margin-top: 0px;margin-bottom: 0px; background-color:#5a91f8; color:white;padding: 0 10px;'>webserv</p>
		</div>
		<ul class="clearfix" style="width: 100%;margin-left: 140px">
			<li><a href="/upload">Upload</a></li>
			<li><a href="/download">Download</a></li>
			<li><a href="/logout">Logout</a></li>
			<li><a href="/close">Close</a></li>
			<li><a href="https://github.com/aoaolion/webserv" target="_blank">About</a></li>
		</ul>
	</header>`
)
