<style type="text/css">
#tab1Div{

	border: red solid 1px;
}
</style>




<div data-dojo-type="dijit/layout/BorderContainer" data-dojo-props="design:'headline', gutters:true, liveSplitters:true">
	<!-- top 区域 -->
	<div data-dojo-type="dijit/layout/ContentPane" data-dojo-props="splitter:false, region:'top'" style="width: 260px;">
    	查询输入区域
    </div>
    <!-- right 区域 -->
    <div data-dojo-type="dijit/layout/ContentPane" data-dojo-props="splitter:true, region:'center'">
    	<div id="roleLoginTable" style="width: 100px; height: 300px;"></div>
    </div>
    <!-- end right -->
</div>