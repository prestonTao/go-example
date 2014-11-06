<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<!-- 暂时禁止浏览器缓存,方便开发 -->
<meta http-equiv="cache-control" content="max-age=0" />
<meta http-equiv="cache-control" content="no-cache" />
<meta http-equiv="expires" content="0" />
<meta http-equiv="expires" content="Tue, 01 Jan 1980 1:00:00 GMT" />
<meta http-equiv="pragma" content="no-cache" />
<title>oss</title>
<!--网页图标-->
<link rel="shortcut icon" href="images/favicon.png" />

<style type="text/css">

@import "static/js/dojo/dijit/themes/claro/claro.css";
/*
@import "dojo/dijit/themes/tundra/tundra.css";
*/
@import "static/js/dojo/dojo/resources/dojo.css";
/*
@import "dojox/grid/resources/tundraGrid.css";
*/
@import "static/js/dojo/gridx/resources/claro/Gridx.css";

html, body {
	width: 100%;
	height: 100%;
	margin: 0;
	overflow:hidden;
}
/* 这个属性不设置大小，页面就不会显示任何东西 */
#rootborder {
	width: 100%;
	height: 100%;
    margin:0;
    padding: 0;
}
</style>

<!-- 
djConfig="parseOnLoad: true"该属性表示在页面加载完成后，启用Dojo的解析模块对Dojo标签进行解析，如果将parseOnLoad:false,声明式小部件就不起作用了。
djConfig="parseOnLoad:true" 表示确定在页面加载完成以后执行解析功能，但解析功能模块的引入要靠 dojo.require("dojo.parser") 来实现。
如果你启用了parseOnLoad: true，碰到页面有这种data-dojo-type，而你又没有事先require了相关的类库，他是会抛出一个js error的。而且说句实话，光是一个dijit.form.ComboBox，他就会require一大堆js文件，真是为这个页面感到费劲啊。
建议将它设置为 false
 -->
<script type="text/javascript" src="static/js/dojo/dojo/dojo.js" data-dojo-config="parseOnLoad: false" charset="utf-8"></script>
<script type="text/javascript" src="static/js/src.js" charset="utf-8"></script>
<script type="text/javascript" src="static/js/main.js" charset="utf-8"></script>

</head>
<body class="claro">

<!-- 页面正在加载的时候，显示的东西 -->
<div id="preloader" style="width:100%;height:100%;background:rgb(246,246,246);padding-top:200px;border:red solid 0px;">
	<div>
		<div style="margin:0 auto; width:450px;font-size:22px;color:rgb(100,100,100);">信都科技支付密码器管理系统(Version 5.1)</div>
		<div style="margin:0 auto; margin-top:10px; width:450px;font-size:16px;color:rgb(250,160,46);">Loading ......</div>
	</div>
</div>


<!-- design属性默认是'headline',当属性值为'headline'时，top区域在最上面。当属性值为'sidebar'时，top区域在右上方 -->
<div data-dojo-type="dijit/layout/BorderContainer" data-dojo-props="design:'headline', gutters:true, liveSplitters:true" id="rootborder">	
	<!-- top 区域 -->
	<div id="indexTop" data-dojo-type="dijit/layout/ContentPane" data-dojo-props="splitter:false, region:'top'" style="width: 260px;">
    	<div style="width:100px;float:left;">选择平台：</div>
        <div style="width:100px;float:left;" id="platFormSelect"></div>
        <div style="width:100px;float:left;">选择服务器：</div>
        <div style="width:100px;float:left;" id="serverSelect"></div>
    </div>
    <!-- end top -->

    <!-- left 区域 -->
    <!-- 这是一个AccordionContainer组件 -->
    <div id="leftTree" data-dojo-type="dijit/layout/AccordionContainer" data-dojo-props="splitter:false, region:'leading'" style="width: 260px;">
    	<!-- 这是Action中的一个Pane -->
    	<div data-dojo-type="dijit/layout/ContentPane" id="ap1"
			data-dojo-props="title: '日志列表'" class="paneAccordion">
            <!-- Pane 中的一个按钮 -->
            <button type="button" data-dojo-type="dijit/form/Button" class="commandButton">
              <span><div class="nvaButton">统计曲线图</div></span>
                <script type="dojo/on" data-dojo-event="click">
                    //这里可以调用通过lang.setObject()添加到方法。
                    index.showTab("统计曲线图","view/tab1.tpl","static/js/tab1.js");
                </script>
            </button><br/>
            <!-- Pane 中的一个按钮 -->
            <button type="button" data-dojo-type="dijit/form/Button" class="commandButton">
              <span><div class="nvaButton">表格</div></span>
                <script type="dojo/on" data-dojo-event="click">
                    //这里可以调用通过lang.setObject()添加到方法。
                    index.showTab("表格","view/gridxdemo.tpl","static/js/gridxdemo.js");
                </script>
            </button><br/>
            <!-- Pane 中的一个按钮 -->
            <button type="button" data-dojo-type="dijit/form/Button" class="commandButton">
              <span><div class="nvaButton">角色登录登出创建和删除</div></span>
                <script type="dojo/on" data-dojo-event="click">
                    //这里可以调用通过lang.setObject()添加到方法。
                    index.showTab("角色登录登出创建和删除","view/roleLogin.tpl","static/js/roleLogin.js");
                </script>
            </button><br/>
            <!-- Pane 中的一个按钮 -->
            <button type="button" data-dojo-type="dijit/form/Button" class="commandButton">
              <span><div class="nvaButton">用户</div></span>
                <script type="dojo/on" data-dojo-event="click">
                    //这里可以调用通过lang.setObject()添加到方法。
                    index.showTab("用户","view/user.tpl","static/js/user.js");
                </script>
            </button><br/>
            <!-- Pane 中的一个按钮 -->
            <button type="button" data-dojo-type="dijit/form/Button" class="commandButton">
              <span><div class="nvaButton">渠道</div></span>
                <script type="dojo/on" data-dojo-event="click">
                    //这里可以调用通过lang.setObject()添加到方法。
                    index.showTab("渠道","view/source.tpl","static/js/source.js");
                </script>
            </button><br/>
            <!-- Pane 中的一个按钮 -->
            <button type="button" data-dojo-type="dijit/form/Button" class="commandButton">
              <span><div class="nvaButton">付费</div></span>
                <script type="dojo/on" data-dojo-event="click">
                    //这里可以调用通过lang.setObject()添加到方法。
                    index.showTab("道具失去和获得","view/gridxdemo.tpl","static/js/gridxdemo.js");
                </script>
            </button><br/>
            <!-- Pane 中的一个按钮 -->
            <button type="button" data-dojo-type="dijit/form/Button" class="commandButton">
              <span><div class="nvaButton">经济</div></span>
                <script type="dojo/on" data-dojo-event="click">
                    //这里可以调用通过lang.setObject()添加到方法。
                    index.showTab("玩家交易和邮件","view/gridxdemo.tpl","static/js/gridxdemo.js");
                </script>
            </button><br/>
            <!-- Pane 中的一个按钮 -->
            <button type="button" data-dojo-type="dijit/form/Button" class="commandButton">
              <span><div class="nvaButton">运营活动</div></span>
                <script type="dojo/on" data-dojo-event="click">
                    //这里可以调用通过lang.setObject()添加到方法。
                    index.showTab("货币变化","view/gridxdemo.tpl","static/js/gridxdemo.js");
                </script>
            </button><br/>
            <!-- Pane 中的一个按钮 -->
            <button type="button" data-dojo-type="dijit/form/Button" class="commandButton">
              <span><div class="nvaButton">预留</div></span>
                <script type="dojo/on" data-dojo-event="click">
                    //这里可以调用通过lang.setObject()添加到方法。
                    index.showTab("公会动态","view/gridxdemo.tpl","static/js/gridxdemo.js");
                </script>
            </button><br/>


        </div>
    </div>
    <!-- end left -->

    <!-- right 区域 -->
    <div id="rightTab" data-dojo-type="dijit/layout/ContentPane" data-dojo-props="splitter:true, region:'center'">
    	<div data-dojo-type="dijit/layout/TabContainer" id="mainTabContainer" data-dojo-props="region: 'center'">
    		<!-- 默认显示的页面 -->
            <div data-dojo-type="dijit/layout/ContentPane" id="tabWelcome" data-dojo-props="title: '概况'">
                <div style="margin:0;padding:0;" data-dojo-type="dijit/layout/BorderContainer" data-dojo-props="design:'headline', gutters:true, liveSplitters:true">
                    <!-- top 区域 
                    <div data-dojo-type="dijit/layout/ContentPane" data-dojo-props="splitter:false, region:'top'" style="width: 260px;">
                        查询输入区域
                    </div>
                    -->
                    <!-- right 区域 -->
                    <div data-dojo-type="dijit/layout/ContentPane" data-dojo-props="splitter:true, region:'center'">
                        <button id="today1" type="button"></button>
                        <button id="today3" type="button"></button>
                        <button id="today7" type="button"></button>
                        <div id="profileGridx" style="width: 400px; height: 300px;"></div>
                        
                        <!-- 整体走向 -->
                        <div data-dojo-type="dijit/TitlePane" data-dojo-props="title: '整体走向'" style="margin-top:20px;">
                            <!-- 注册/登录 -->
                            <button id="profile_trend_bt_login" type="button"></button>
                            <!-- acu/pcu -->
                            <button id="profile_trend_bt_acu" type="button"></button>
                            <!-- 付费人数/次数 -->
                            <button id="profile_trend_bt_prepaidCount" type="button"></button>
                            <!-- 充值/消费 -->
                            <button id="profile_trend_bt_prepaid" type="button"></button>
                            <!-- 整体流失 -->
                            <button id="profile_trend_bt_outflow" type="button"></button>
                            <!-- 七日留存 -->
                            <button id="profile_trend_bt_today7" type="button"></button>

                            <div id="chartOne" style="width: 1000px; height: 340px;"></div>
                        </div>

                        <!-- 时段分析 -->
                        <div data-dojo-type="dijit/TitlePane" data-dojo-props="title: '时段分析'" style="margin-top:20px;">
                            <!-- 注册 -->
                            <button id="profile_period_bt_regist" type="button"></button>
                            <!-- 登录 -->
                            <button id="profile_period_bt_login" type="button"></button>
                            <!-- 在线 -->
                            <button id="profile_period_bt_online" type="button"></button>
                            <!-- 付费人数 -->
                            <button id="profile_period_bt_prepaidPersonCount" type="button"></button>
                            <!-- 付费次数 -->
                            <button id="profile_period_bt_prepaidCount" type="button"></button>
                            <!-- 充值 -->
                            <button id="profile_period_bt_prepaid" type="button"></button>
                            <!-- 消费 -->
                            <button id="profile_period_bt_consumption" type="button"></button>

                            <div style="width:500px;margin:0 auto;font-size: 30px;">各时段注册用户与登陆用户对比</div>
                            <div id="periodAnalyzeChart" style="width: 1000px; height: 240px;"></div>
                        </div>
                        
                    </div>
                    <!-- end right -->
                </div>
            </div>
        </div>
    </div>
    <!-- end right -->
</div>

</body>
</html>
