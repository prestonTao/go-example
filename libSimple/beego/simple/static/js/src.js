require([
	"dojo/ready",
	"dojo/parser",
	"dojo/dom-style",
	"dojo/_base/fx",
	"dojo/_base/lang",
	"dijit/registry",
	"dijit/layout/ContentPane",
	"dojo/when",
	"dojo/io/script",
	"dijit/form/Select"
], function(ready,parser,domStyle,fx,lang,registry,ContentPane,when,ioScript, Select){


ready(function(){
	loadPage();
	buildTop();
});



/**
 * top区域
 **/
function buildTop(){
	var platFormSelect = new Select({
        name: "platFormSelect",
        options: [
            { label: "多玩", value: "Tennessee", selected: true },
            { label: "YY", value: "Virginia" },
            { label: "WA", value: "Washington" },
            { label: "FL", value: "Florida" },
            { label: "CA", value: "California" }
        ]
    }).placeAt("platFormSelect");

    platFormSelect.on("change", function(){
		//alert(1);
	});

	var serverSelect = new Select({
        name: "serverSelect",
        options: [
            { label: "服务器一", value: "Tennessee" },
            { label: "服务器二", value: "Virginia", selected: true },
            { label: "服务器三", value: "Washington" },
            { label: "服务器四", value: "Florida" },
            { label: "服务器五", value: "California" }
        ]
    }).placeAt("serverSelect");

    serverSelect.on("change", function(){
		//alert(2);
	});

}




/**
 * 全局方法，整个页面任意位置都能调用
 * tabName: tab显示的title名称
 * htmlUrl: 加载的html页面路径
 * jsUrl：加载的js文件路径，可以为空
 * cssUrl：加载的css样式文件路径，可以为空
 **/
lang.setObject("index.showTab",function(tabName,htmlUrl,jsUrl){
	var tabContainer = registry.byId("mainTabContainer");
	//tab的id直接用htmlUrl路径命名把
	var tabOne = registry.byId(htmlUrl);
	if(typeof tabOne === "undefined"){
		tabOne = new ContentPane({
			id:htmlUrl,
			title:tabName,
			closable:true,  //是否显示关闭按钮
			//content: "We are known for our drinks."
			href:htmlUrl
		});
		
		when(tabContainer.addChild(tabOne),
			when(tabContainer.selectChild(tabOne),function(){
				if(jsUrl != null && jsUrl != ""){
					ioScript.get({url:jsUrl});
				}
			})
		);
	}else{
		//选中tab
		tabContainer.selectChild(tabOne);
	}
});



function loadPage(){
	parser.parse().then(function(objects){
		fx.fadeOut({  //Get rid of the loader once parsing is done
			node: "preloader",
			duration: 0,
			onEnd: function() {
				domStyle.set("preloader","display","none");
			}
		}).play();
	});
}





});