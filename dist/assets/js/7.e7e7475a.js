(window.webpackJsonp=window.webpackJsonp||[]).push([[7],{200:function(t,e,s){"use strict";s.r(e);var n=s(0),i=Object(n.a)({},function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"content"},[t._m(0),t._v(" "),t._m(1),t._v(" "),s("p",[t._v("​\tWelcome to the SingerAPI, the "),s("a",{attrs:{href:"https://en.wikipedia.org/wiki/I_Am_a_Singer_(Chinese_TV_series)",target:"_blank",rel:"noopener noreferrer"}},[t._v("I-Am-a-Singer"),s("OutboundLink")],1),t._v(" API! This is an api for the singers and songs information from Chinese TV series "),s("a",{attrs:{href:"https://en.wikipedia.org/wiki/I_Am_a_Singer_(Chinese_TV_series)",target:"_blank",rel:"noopener noreferrer"}},[t._v("I-Am-a-Singer"),s("OutboundLink")],1),t._v(". Our project provides simple HTTP requests to access the related resources.")]),t._v(" "),t._m(2),t._v(" "),t._m(3),t._v(" "),t._m(4),t._v(" "),s("p",[t._v("​\tBefore type in our first eample, please ensure you have follow the installation instructions of SingerAPI api resource server and run server in local.")]),t._v(" "),s("p",[t._v("​\tThen, let's make our first API request to the SingerAPI with "),s("a",{attrs:{href:"https://github.com/jakubroztocil/httpie/",target:"_blank",rel:"noopener noreferrer"}},[t._v("httpie"),s("OutboundLink")],1),t._v("!")]),t._v(" "),t._m(5),s("p",[t._v("​\tThen you will get response that:")]),t._v(" "),t._m(6),s("p",[t._v("​\tAlso you can open SingerAPI home site with server running in the local : http://127.0.0.1:8080/singerapi/")]),t._v(" "),t._m(7),t._v(" "),s("p",[t._v("The Base URL is the root URL for all of the API, if you ever make a request to swapi and you get back a 404 NOT FOUND response then check the Base URL first.")]),t._v(" "),s("p",[t._v("The Base URL for SingerAPI is:")]),t._v(" "),s("p",[t._v("http://127.0.0.1:8080/singerapi/api")]),t._v(" "),s("p",[t._v("The documentation below assumes you are prepending the Base URL to the endpoints in order to make requests.")]),t._v(" "),t._m(8),t._v(" "),s("p",[t._v("Swapi has rate limiting to prevent malicious abuse (as if anyone would abuse Star Wars data!) and to make sure our service can handle a potentially large amount of traffic. Rate limiting is done via IP address and is currently limited to 10,000 API request per day. This is enough to request all the data on the website at least ten times over. There should be no reason for hitting the rate limit.")]),t._v(" "),t._m(9),t._v(" "),s("p",[t._v("Swapi is a completely open API. No authentication is required to query and get data. This also means that we've limited what you can do to just GET-ing the data. If you find a mistake in the data, then tweet the author or email him.")]),t._v(" "),s("p",[t._v("JSON Schema\nAll resources support JSON Schema. Making a request to /api/resource/schema will give you the details of that resource. This will allow you to programmatically inspect the attributes of that resource and their types.")]),t._v(" "),t._m(10),t._v(" "),s("p",[t._v("All resources support a search parameter that filters the set of resources returned. This allows you to make queries like:")]),t._v(" "),s("p",[t._v("http://127.0.0.1/singerapi/api/singers/?name=xxx")]),t._v(" "),s("p",[t._v("All searches will use case-insensitive partial matches on the set of search fields. To see the set of search fields for each resource, check out the individual resource documentation. For more information on advanced search terms see here.")]),t._v(" "),t._m(11),t._v(" "),s("p",[t._v("SWAPI provides two encodings for you to render the data with:")]),t._v(" "),t._m(12),t._v(" "),s("p",[t._v("JSON is the standard data format provided by SingerAPI by default.")]),t._v(" "),t._m(13),t._v(" "),t._m(14),t._v(" "),s("p",[t._v("The API resource provides information on all available resources within the API.")]),t._v(" "),t._m(15),t._v(" "),t._m(16),t._m(17),t._v(" "),t._m(18),t._m(19),t._v(" "),t._m(20),t._v(" "),s("br"),t._v(" "),t._m(21),t._v(" "),t._m(22),t._v(" "),t._m(23),t._v(" "),t._m(24),t._v(" "),t._m(25),t._v(" "),t._m(26),t._m(27),t._v(" "),t._m(28),t._m(29),t._v(" "),t._m(30),t._v(" "),s("br"),t._v(" "),t._m(31),t._v(" "),t._m(32),t._v(" "),t._m(33),t._v(" "),t._m(34),t._v(" "),t._m(35),t._v(" "),t._m(36),t._m(37),t._v(" "),t._m(38),t._m(39),t._v(" "),t._m(40),t._v(" "),s("br"),t._v(" "),t._m(41),t._v(" "),t._m(42),t._v(" "),t._m(43),t._v(" "),t._m(44),t._v(" "),t._m(45),t._v(" "),t._m(46),t._m(47),t._v(" "),t._m(48),t._m(49),t._v(" "),t._m(50),t._v(" "),s("br"),t._v(" "),t._m(51),t._v(" "),t._m(52),t._v(" "),t._m(53),t._v(" "),t._m(54),t._v(" "),t._m(55),t._v(" "),t._m(56),t._m(57),t._v(" "),t._m(58),t._m(59),t._v(" "),t._m(60),t._v(" "),s("br"),t._v(" "),t._m(61),t._v(" "),t._m(62),t._v(" "),t._m(63),t._v(" "),t._m(64),t._v(" "),t._m(65),t._v(" "),t._m(66),t._m(67),t._v(" "),t._m(68),t._m(69),t._v(" "),t._m(70),t._v(" "),t._m(71),t._v(" "),t._m(72),t._v(" "),s("br"),t._v(" "),t._m(73),t._v(" "),t._m(74),t._v(" "),t._m(75),t._v(" "),t._m(76),t._v(" "),t._m(77),t._v(" "),t._m(78),t._m(79),t._v(" "),t._m(80),t._m(81),t._v(" "),t._m(82),t._v(" "),t._m(83),t._v(" "),t._m(84),t._v(" "),s("br"),t._v(" "),t._m(85)])},[function(){var t=this.$createElement,e=this._self._c||t;return e("h1",{attrs:{id:"documentation"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#documentation","aria-hidden":"true"}},[this._v("#")]),this._v(" Documentation")])},function(){var t=this.$createElement,e=this._self._c||t;return e("h2",{attrs:{id:"introduction"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#introduction","aria-hidden":"true"}},[this._v("#")]),this._v(" Introduction")])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[this._v("​\tTo get started, grab a cup of your favorite beverage and read following "),e("strong",[this._v("Getting Started")]),this._v(" section.")])},function(){var t=this.$createElement,e=this._self._c||t;return e("h2",{attrs:{id:"getting-started"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#getting-started","aria-hidden":"true"}},[this._v("#")]),this._v(" Getting started")])},function(){var t=this.$createElement,e=this._self._c||t;return e("blockquote",[e("p",[this._v("TODO: complete installation instruction of api resource server.")]),this._v(" "),e("p",[this._v("TODO: complete HTTP response result.")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"language-shell extra-class"},[e("pre",{pre:!0,attrs:{class:"language-text"}},[e("code",[this._v("http http://127.0.0.1:8080/singerapi/api/seasons/\n")])])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"language-shell extra-class"},[e("pre",{pre:!0,attrs:{class:"language-text"}},[e("code",[this._v("...\n")])])])},function(){var t=this.$createElement,e=this._self._c||t;return e("h2",{attrs:{id:"base-url"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#base-url","aria-hidden":"true"}},[this._v("#")]),this._v(" Base URL")])},function(){var t=this.$createElement,e=this._self._c||t;return e("h2",{attrs:{id:"rate-limiting"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#rate-limiting","aria-hidden":"true"}},[this._v("#")]),this._v(" Rate limiting")])},function(){var t=this.$createElement,e=this._self._c||t;return e("h2",{attrs:{id:"authentication"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#authentication","aria-hidden":"true"}},[this._v("#")]),this._v(" Authentication")])},function(){var t=this.$createElement,e=this._self._c||t;return e("h2",{attrs:{id:"searching"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#searching","aria-hidden":"true"}},[this._v("#")]),this._v(" Searching")])},function(){var t=this.$createElement,e=this._self._c||t;return e("h2",{attrs:{id:"encodings"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#encodings","aria-hidden":"true"}},[this._v("#")]),this._v(" Encodings")])},function(){var t=this.$createElement,e=this._self._c||t;return e("h3",{attrs:{id:"json"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#json","aria-hidden":"true"}},[this._v("#")]),this._v(" JSON")])},function(){var t=this.$createElement,e=this._self._c||t;return e("h2",{attrs:{id:"resources"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#resources","aria-hidden":"true"}},[this._v("#")]),this._v(" Resources")])},function(){var t=this.$createElement,e=this._self._c||t;return e("h3",{attrs:{id:"api"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#api","aria-hidden":"true"}},[this._v("#")]),this._v(" API")])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Example request:")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"language-http extra-class"},[e("pre",{pre:!0,attrs:{class:"language-http"}},[e("code",[this._v("http 127.0.0.1:8080/singerapi/api\n")])])])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Example response:")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"language-http extra-class"},[e("pre",{pre:!0,attrs:{class:"language-http"}},[e("code",[this._v("\n\n"),e("span",{attrs:{class:"token response-status"}},[this._v("HTTP/1.0 "),e("span",{attrs:{class:"token property"}},[this._v("200 OK")])]),this._v("\n"),e("span",{attrs:{class:"token header-name keyword"}},[this._v("Content-Type:")]),this._v(' application/json\n{\n    "all seasons": "http://127.0.0.1/singerapi/api/seasons/",\n    "specific season": "http://127.0.0.1/singerapi/api/seasons/{id}/",\n    "singers in season": "http://127.0.0.1/singerapi/api/seasons/{id}/singers",\n    "songs in season": "http://127.0.0.1/singerapi/api/seasons/{id}/songs",\n    "specific singer": "http://127.0.0.1/singerapi/api/singers/?name={name}",\n    "specific song": "http://127.0.0.1/singerapi/api/songs/?name={name}"\n}'),e("span",{attrs:{class:"token application/json"}},[this._v("\n\n")])])])])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Attributes:")])])},function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("ul",[s("li",[s("code",[t._v("all seasons")]),t._v(" string -- The URL root for all seasons resources")]),t._v(" "),s("li",[s("code",[t._v("specific seasons")]),t._v(" string -- The URL root for specific season resources")]),t._v(" "),s("li",[s("code",[t._v("singers in seasons")]),t._v(" string -- The URL root for all singers in specific season")]),t._v(" "),s("li",[s("code",[t._v("songs in seansons")]),t._v(" string -- The URL root for all songs in specific season")]),t._v(" "),s("li",[s("code",[t._v("specific singer")]),t._v(" string -- The URL root for specific singer resources")]),t._v(" "),s("li",[s("code",[t._v("specific song")]),t._v(" string -- The URL root for specific song resources")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("h3",{attrs:{id:"all-seasons"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#all-seasons","aria-hidden":"true"}},[this._v("#")]),this._v(" All seasons")])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[this._v("All URL to specific seasons in "),e("em",[this._v("I Am a Singer")]),this._v(" .")])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Endpoints")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("ul",[e("li",[e("code",[this._v("/seasons/")]),this._v(" -- get all the season resources")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Example request:")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"language-http extra-class"},[e("pre",{pre:!0,attrs:{class:"language-http"}},[e("code",[this._v("http http://127.0.0.1:8080/singerapi/api/seasons/\n")])])])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Example response:")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"language-http extra-class"},[e("pre",{pre:!0,attrs:{class:"language-http"}},[e("code",[e("span",{attrs:{class:"token response-status"}},[this._v("HTTP/1.0 "),e("span",{attrs:{class:"token property"}},[this._v("200 OK")])]),this._v("\n"),e("span",{attrs:{class:"token header-name keyword"}},[this._v("Content-Type:")]),this._v(' application/json\n{\n\t"seasons" : {\n\t\t"我是歌手第一季" : "http://127.0.0.1:8080/singerapi/api/seasons/1/",\n\t\t"我是歌手第一季" : "http://127.0.0.1:8080/singerapi/api/seasons/2/",\n\t\t"我是歌手第一季" : "http://127.0.0.1:8080/singerapi/api/seasons/3/",\n\t\t"我是歌手第一季" : "http://127.0.0.1:8080/singerapi/api/seasons/4/",\n\t\t"我是歌手第一季(歌手2017)" : "http://127.0.0.1:8080/singerapi/api/seasons/5/",\n\t\t"我是歌手第一季(歌手2018)" : "http://127.0.0.1:8080/singerapi/api/seasons/6/"\n\t}\n\t"url" : "http://127.0.0.1/singerapi/api/seasons/"\n}\n')])])])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Attributes:")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("ul",[e("li",[e("code",[this._v("seasons")]),this._v(" json -- All seasons and corresponding URL in "),e("em",[this._v("I Am a Singer")]),this._v(".")]),this._v(" "),e("li",[e("code",[this._v("url")]),this._v(" string -- The URL of this resource.")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("h3",{attrs:{id:"specific-season"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#specific-season","aria-hidden":"true"}},[this._v("#")]),this._v(" Specific season")])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[this._v("A Seasons resource is a specific "),e("em",[this._v("I Am a Singer")]),this._v(" season.")])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Endpoints")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("ul",[e("li",[e("code",[this._v("/seasons/{id}/")]),this._v(" -- get a specific season resource")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Example request:")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"language-http extra-class"},[e("pre",{pre:!0,attrs:{class:"language-http"}},[e("code",[this._v("http http://127.0.0.1:8080/singerapi/api/seasons/1/\n")])])])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Example response:")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"language-http extra-class"},[e("pre",{pre:!0,attrs:{class:"language-http"}},[e("code",[e("span",{attrs:{class:"token response-status"}},[this._v("HTTP/1.0 "),e("span",{attrs:{class:"token property"}},[this._v("200 OK")])]),this._v("\n"),e("span",{attrs:{class:"token header-name keyword"}},[this._v("Content-Type:")]),this._v(' application/json\n{\n\t"title" : "我是歌手第一季",\n\t"broadcast time" : "2013.1.18-2013.4.12",\n\t"hosts" : "胡海泉/陈羽凡/沙宝亮/何炅/汪涵",\n\t"winner" : "羽泉",\n\t"singers" : "http://127.0.0.1/singerapi/api/seasons/1/singers/",\n\t"songs" : "http://127.0.0.1/singerapi/api/seasons/1/songs/",\n\t"url" : "http://127.0.0.1/singerapi/api/seasons/1/"\n}\n')])])])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Attributes:")])])},function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("ul",[s("li",[s("code",[t._v("title")]),t._v(" string -- The season title.")]),t._v(" "),s("li",[s("code",[t._v("broadcast time")]),t._v(" string -- The broadcasts time of this season.")]),t._v(" "),s("li",[s("code",[t._v("hosts")]),t._v(" string -- The hosts of this season.")]),t._v(" "),s("li",[s("code",[t._v("winner")]),t._v(" string -- The finals winner of this season.")]),t._v(" "),s("li",[s("code",[t._v("singers")]),t._v(" string -- The URL of singers joined to this season.")]),t._v(" "),s("li",[s("code",[t._v("songs")]),t._v(" string -- The URL of songs performed in thin season.")]),t._v(" "),s("li",[s("code",[t._v("url")]),t._v(" string -- The URL of this resource.")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("h3",{attrs:{id:"singers-in-specific-season"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#singers-in-specific-season","aria-hidden":"true"}},[this._v("#")]),this._v(" Singers in specific season")])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[this._v("All singers performed in specific season in "),e("em",[this._v("I Am a Singer")]),this._v(" .")])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Endpoints")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("ul",[e("li",[e("code",[this._v("/seasons/{id}/singers/")]),this._v(" -- get all singers contributed in specific season")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Example request:")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"language-http extra-class"},[e("pre",{pre:!0,attrs:{class:"language-http"}},[e("code",[this._v("http http://127.0.0.1:8080/singerapi/api/seasons/1/singers/\n")])])])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Example response:")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"language-http extra-class"},[e("pre",{pre:!0,attrs:{class:"language-http"}},[e("code",[e("span",{attrs:{class:"token response-status"}},[this._v("HTTP/1.0 "),e("span",{attrs:{class:"token property"}},[this._v("200 OK")])]),this._v("\n"),e("span",{attrs:{class:"token header-name keyword"}},[this._v("Content-Type:")]),this._v(' application/json\n{\n\t"season" : "我是歌手第一季"\n\t"singers" : {\n\t\t"羽泉" : "http://127.0.0.1:8080/singerapi/api/singers/?singer=羽泉",\n\t\t"杨宗纬" : "http://127.0.0.1:8080/singerapi/api/singers/?singer=杨宗纬",\n\t\t"林志炫" : "http://127.0.0.1:8080/singerapi/api/singers/?singer=林志炫"\n\t\t...\n\t}\n\t"url" : "http://127.0.0.1:8080/singerapi/api/seasons/1/singers/"\n}\n')])])])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Attributes:")])])},function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("ul",[s("li",[s("code",[t._v("season")]),t._v(" string -- The name of this season.")]),t._v(" "),s("li",[s("code",[t._v("singers")]),t._v(" json -- All singers in thin season and corresponding URL.")]),t._v(" "),s("li",[s("code",[t._v("url")]),t._v(" string -- The URL of this resource.")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("h3",{attrs:{id:"songs-in-specific-season"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#songs-in-specific-season","aria-hidden":"true"}},[this._v("#")]),this._v(" Songs in specific season")])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[this._v("All songs performed in specific season in "),e("em",[this._v("I Am a Singer")]),this._v(" .")])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Endpoints")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("ul",[e("li",[e("code",[this._v("/seasons/{id}/songs/")]),this._v(" -- get all songs performed in specific season")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Example request:")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"language-http extra-class"},[e("pre",{pre:!0,attrs:{class:"language-http"}},[e("code",[this._v("http http://127.0.0.1:8080/singerapi/api/seasons/1/songs/\n")])])])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Example response:")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"language-http extra-class"},[e("pre",{pre:!0,attrs:{class:"language-http"}},[e("code",[e("span",{attrs:{class:"token response-status"}},[this._v("HTTP/1.0 "),e("span",{attrs:{class:"token property"}},[this._v("200 OK")])]),this._v("\n"),e("span",{attrs:{class:"token header-name keyword"}},[this._v("Content-Type:")]),this._v(' application/json\n{\n\t"season" : "我是歌手第一季",\n\t"songs" : {\n\t\t"浮夸" : "http://127.0.0.1:8080/singerapi/api/songs/?song=浮夸",\n\t\t"相见恨晚" : "http://127.0.0.1:8080/singerapi/api/songs/?song=相见恨晚",\n\t\t"往事随风" : "http://127.0.0.1:8080/singerapi/api/songs/?song=往事随风",\n\t\t...\n\t}\n\t"url" : "http://127.0.0.1:8080/singerapi/api/seasons/1/songs/"\n}\n')])])])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Attributes:")])])},function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("ul",[s("li",[s("code",[t._v("season")]),t._v(" string -- The name of this season.")]),t._v(" "),s("li",[s("code",[t._v("songs")]),t._v(" json -- All songs performed in this season and correspongding URL.")]),t._v(" "),s("li",[s("code",[t._v("url")]),t._v(" string -- The URL of this resource.")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("h3",{attrs:{id:"specific-singer"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#specific-singer","aria-hidden":"true"}},[this._v("#")]),this._v(" Specific singer")])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[this._v("A Singer resource is a "),e("em",[this._v("I Am a Singer")]),this._v(" participated singer.")])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Endpoints")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("ul",[e("li",[e("code",[this._v("/singers/?singer={name}")]),this._v(" -- get a specific singers resource")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Example request:")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"language-http extra-class"},[e("pre",{pre:!0,attrs:{class:"language-http"}},[e("code",[this._v("http http://127.0.0.1:8080/singerapi/api/singers/?singer=张韶涵\n")])])])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Example response:")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"language-http extra-class"},[e("pre",{pre:!0,attrs:{class:"language-http"}},[e("code",[e("span",{attrs:{class:"token response-status"}},[this._v("HTTP/1.0 "),e("span",{attrs:{class:"token property"}},[this._v("200 OK")])]),this._v("\n"),e("span",{attrs:{class:"token header-name keyword"}},[this._v("Content-Type:")]),this._v(' application/json\n{\n\t"singer" : "张韶涵",\n\t"songs": {\n\t\t"阿刁" : "http://127.0.0.1:8080/singerapi/api/songs/?name=阿刁",\n\t\t"梦里花" : "http://127.0.0.1:8080/singerapi/api/songs/?name=梦里花",\n\t\t...\n\t},\n\t"url" : "http://127.0.0.1/singerapi/api/singers/?name=张韶涵"\n}'),e("span",{attrs:{class:"token application/json"}},[this._v("\n\n")])])])])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Attributes:")])])},function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("ul",[s("li",[s("code",[t._v("singer")]),t._v(" string -- The singer name.")]),t._v(" "),s("li",[s("code",[t._v("songs")]),t._v(" json-- All singer's contributed songs and corresponding URL.")]),t._v(" "),s("li",[s("code",[t._v("url")]),t._v(" string -- The URL of this resource.")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Search Fields:")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("ul",[e("li",[e("code",[this._v("singer")])])])},function(){var t=this.$createElement,e=this._self._c||t;return e("h3",{attrs:{id:"specific-song"}},[e("a",{staticClass:"header-anchor",attrs:{href:"#specific-song","aria-hidden":"true"}},[this._v("#")]),this._v(" Specific song")])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[this._v("A Song resource is a song performed in "),e("em",[this._v("I Am a Singer")]),this._v(" .")])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Endpoints")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("ul",[e("li",[e("code",[this._v("/songs/?song={name}")]),this._v(" -- get a specific songs resource")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Example request:")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"language-http extra-class"},[e("pre",{pre:!0,attrs:{class:"language-http"}},[e("code",[this._v("http http://127.0.0.1:8080/singerapi/api/songs/?song=夜夜夜夜\n")])])])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Example response:")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"language-http extra-class"},[e("pre",{pre:!0,attrs:{class:"language-http"}},[e("code",[e("span",{attrs:{class:"token response-status"}},[this._v("HTTP/1.0 "),e("span",{attrs:{class:"token property"}},[this._v("200 OK")])]),this._v("\n"),e("span",{attrs:{class:"token header-name keyword"}},[this._v("Content-Type:")]),this._v(' application/json\n{\n\t"song" : "夜夜夜夜",\n\t"singer" : "林志炫",\n\t"duration" : "05:00"\n\t"album" : "我是歌手第一季 第10期",\n\t"season" : "我是歌手第一季"\n\t"url" : "http://127.0.0.1/singerapi/api/songs/?song=夜夜夜夜"\n}'),e("span",{attrs:{class:"token application/json"}},[this._v("\n\n")])])])])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Attributes:")])])},function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("ul",[s("li",[s("code",[t._v("song")]),t._v(" string -- The name of the specific song")]),t._v(" "),s("li",[s("code",[t._v("singer")]),t._v(" string -- The name of singer who performed the song in "),s("em",[t._v("I Am a Singer")]),t._v(".")]),t._v(" "),s("li",[s("code",[t._v("duration")]),t._v(" string -- The time of performing the song.")]),t._v(" "),s("li",[s("code",[t._v("album")]),t._v(" string -- The name of the album which the song released in.")]),t._v(" "),s("li",[s("code",[t._v("season")]),t._v(" string -- The season in which the song was performed in "),s("em",[t._v("I Am a Singer")]),t._v(".")]),t._v(" "),s("li",[s("code",[t._v("url")]),t._v(" string -- The URL of this resource.")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("p",[e("strong",[this._v("Search Fields:")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("ul",[e("li",[e("code",[this._v("song")])])])},function(){var t=this.$createElement,e=this._self._c||t;return e("blockquote",[e("p",[this._v("refs: https://swapi.co/documentation")])])}],!1,null,null,null);i.options.__file="README.md";e.default=i.exports}}]);