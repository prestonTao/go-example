require(['dojo/store/Memory',
	"dojox/charting/Chart",
	"dojox/charting/axis2d/Default",
    "dojox/charting/plot2d/StackedLines",
    "dojox/charting/plot2d/Lines",
    "dojox/charting/plot2d/Areas",
    "dojox/charting/plot2d/Bars",
    "dojox/charting/plot2d/Columns",
    "dojox/charting/plot2d/Grid",
    "dojox/charting/plot2d/Bubble",
    "dojox/charting/plot2d/Indicator",
    "dojox/charting/themes/Claro" ,
    "dojo/ready",
    'gridx/Grid',  
    'gridx/core/model/cache/Sync',
    'gridx/core/model/cache/Async',
    "dojo/store/JsonRest",
    'gridx/allModules',
    'gridx/modules/SingleSort',
    "dijit/form/Button",
    "dojox/charting/action2d/Tooltip"],
  function(Memory, Chart, Default, Lines, Areas, StackedLines, Bars, Columns, Grid2d, Bubble, Indicator, Wetland, ready,  Grid, Sync, Async, JsonRestStore, modules, SingleSort,
    Button, Tooltip){


ready(function(){
	buildGridx();
	buildTrend();
    buildPeriod();
});


//创建表格
function buildGridx(){

    

	new Button({
        label: "今日数据",
        onClick: function(){
            // Do something:
            // dom.byId("result1").innerHTML += "Thank you! ";
        }
    }, "today1");

    new Button({
        label: "三日数据",
        onClick: function(){
            // Do something:
            // dom.byId("result1").innerHTML += "Thank you! ";
        }
    }, "today3");

    new Button({
        label: "七日数据",
        onClick: function(){
            // Do something:
            // dom.byId("result1").innerHTML += "Thank you! ";
        }
    }, "today7");

	var items = [{"id": 1, "data":"2014-4-21", "regist":"800000 <label style='color:red;'>↓</label>20%", 
					"login":"30000 <label style='color:red;'>↓</label>50%", "prepaidCount":"150 <label style='color:green;'>↑</label>50%", "arpu":"",
					"acu":"", "pcu":"", "prepaid":"", "spent":"222+30000"},
				{"id": 2, "data":"2014-4-20", "regist":"", "login":"", "prepaidCount":"", "arpu":"",
					"acu":"", "pcu":"", "prepaid":"", "spent":""}];

	var store = new Memory({data: items});
	//定义列结构
    var columns = [  
        {field: 'data', name: '日期'},
        {field: 'regist', name: '注册'},
        {field: 'login', name: '登录'},
        {field: 'prepaidCount', name: '付费人数'},
        {field: 'arpu', name: 'arpu'},
        {field: 'acu', name: 'acu'},
        {field: 'pcu', name: 'pcu'},
        {field: 'prepaid', name: '总充值'},
        {field: 'spent', name: '总消费'}
    ]; 

    var grid = new Grid({
        cacheClass: Sync,
        store: store,
        structure: columns,
        autoHeight: true,
        pageSize:2 //每页显示多少条数据
        
    }, 'profileGridx');
    grid.startup();
}

//整体走势图区域
function buildTrend(){
    new Button({label: "三日数据",
        onClick: function(){
            // Do something:
            // dom.byId("result1").innerHTML += "Thank you! ";
        }
    }, "profile_trend_bt_login");
    new Button({label: "acu/pcu",
        onClick: function(){
            // Do something:
            // dom.byId("result1").innerHTML += "Thank you! ";
        }
    }, "profile_trend_bt_acu");
    new Button({label: "付费人数/次数",
        onClick: function(){
            // Do something:
            // dom.byId("result1").innerHTML += "Thank you! ";
        }
    }, "profile_trend_bt_prepaidCount");
    new Button({label: "充值/消费",
        onClick: function(){
            // Do something:
            // dom.byId("result1").innerHTML += "Thank you! ";
        }
    }, "profile_trend_bt_prepaid");
    new Button({label: "整体流失",
        onClick: function(){
            // Do something:
            // dom.byId("result1").innerHTML += "Thank you! ";
        }
    }, "profile_trend_bt_outflow");
    new Button({label: "七日留存",
        onClick: function(){
            // Do something:
            // dom.byId("result1").innerHTML += "Thank you! ";
        }
    }, "profile_trend_bt_today7");



	var c = new Chart("chartOne",{
        title: "ACU/PCU数据对比",
        titlePos: "top",
        titleGap: 25,
        titleFont: "normal normal normal 15pt Arial",
        titleFontColor: "orange"
    });
	c.addPlot("default")
    .addPlot("default", {type: StackedLines,
        //labelOffset: -30,
        //shadows: { dx:2, dy:2, dw:2 },
        // tension: "S",
        tension:3,
        markers: true //显示线上的点
    })
    // .addPlot("default", { type: Grid2d,
    //      hMajorLines: true,
    //      hMinorLines: false,
    //      vMajorLines: true,
    //      vMinorLines: false,
    //      majorHLine: { color: "green", width: 3 },
    //      majorVLine: { color: "red", width: 3 }
    // })
    .addPlot("threshold", { type: Indicator,
        vertical: false,
        lineStroke: { color: "red", style: "ShortDash"},
        stroke: null,
        outline: null,
        fill: null,
        offset: { y: -7, x: -10 },
        values: 15
    })
    .addPlot("threshold", { type: Indicator,
        vertical: false,
        lineStroke: { color: "red", style: "ShortDash"},
        stroke: null,
        outline: null,
        fill: null,
        offset: { y: -7, x: -10 },
        values: 14
    })
    .addPlot("threshold", { type: Indicator,
        vertical: false,
        lineStroke: { color: "red", style: "ShortDash"},
        labels: "none",
        values: 15
    })
    .addPlot("threshold", { type: Indicator,
        lineStroke: { color: "red", style: "ShortDash"},
        labels: "none",
        values: 15
    })
    .addSeries("markers", [ 8, 17, 30 ], { plot: "threshold" })
    // , {type: StackedLines,
    //     //labelOffset: -30,
    //     //shadows: { dx:2, dy:2, dw:2 },
    //     // tension: "S",
    //     tension:3,
    //     markers: true //显示线上的点
    // }
    // .addPlot("default", {type: Grid2d, enableCache:true, renderOnAxis: false})
    // .addPlot("default", { type: Bars, gap: 5, minBarSize: 3, maxBarSize: 20 })
    // .addPlot("default", {type: Columns, enableCache: true})
    // .addPlot("default", { type: Bars, hAxis: "cool x", vAxis: "super y" })
    // .addPlot("default", { type: Bubble, fill: "red" })
	.addAxis("x", {fixLower: "major", fixUpper: "mijor",
        majorLabels: true,
        minorTicks: true,
        minorLabels: true,
        microTicks: false,
        majorTickStep: 1,
        minorTickStep: 1,
        microTickStep: 1,
        labels: [
            {value: 0, text: ""},
            {value: 1, text: "2014.1.1"},
            {value: 2, text: "2014.1.2"},
            {value: 3, text: "2014.1.3"},
            {value: 4, text: "2014.1.4"},
            {value: 5, text: "2014.1.5"},
            {value: 6, text: "2014.1.6"},
            {value: 7, text: "2014.1.7"},
            {value: 8, text: "2014.1.8"},
            {value: 9, text: "2014.1.9"},
            {value: 10, text: "2014.1.10"},
            {value: 11, text: "2014.1.11"},
            {value: 12, text: "2014.1.12"},
            {value: 13, text: "2014.1.13"},
            {value: 14, text: "2014.1.14"},
            {value: 15, text: "2014.1.15"},
            {value: 16, text: "2014.1.16"}
        ]
    })
    .addAxis("y", {vertical: true, fixLower: "major", fixUpper: "mijor", min: 0})
    .addAxis("other y", {vertical: true, leftBottom: false})
    .setTheme(Wetland)
    .addSeries("Series A", [1, 2, 0.5, 1.5, 1, 2.8, 0.4, 5, 7, 3, 2, 7, 2, 3, 2, 3])
    .addSeries("Series B", [2.6, 1.8, 2, 1, 1.4, 0.7, 2, 4, 3, 2, 6, 7, 2, 8,3, 4])
    // .addSeries("Series C", [6.3, 1.8, 3, 0.5, 4.4, 2.7, 2])
    .render();

    // var tip = new Tooltip(c, "default");
    new Tooltip(c, "default", {
       text: function(o){
          return "Element at index: "+o.index;
       }
    });
}


//构建时段分析区域
function buildPeriod(){

    new Button({label: "注册",
        onClick: function(){
            // Do something:
            // dom.byId("result1").innerHTML += "Thank you! ";
        }
    }, "profile_period_bt_regist");
    new Button({label: "登录",
        onClick: function(){
            // Do something:
            // dom.byId("result1").innerHTML += "Thank you! ";
        }
    }, "profile_period_bt_login");
    new Button({label: "在线",
        onClick: function(){
            // Do something:
            // dom.byId("result1").innerHTML += "Thank you! ";
        }
    }, "profile_period_bt_online");
    new Button({label: "付费人数",
        onClick: function(){
            // Do something:
            // dom.byId("result1").innerHTML += "Thank you! ";
        }
    }, "profile_period_bt_prepaidPersonCount");
    new Button({label: "付费次数",
        onClick: function(){
            // Do something:
            // dom.byId("result1").innerHTML += "Thank you! ";
        }
    }, "profile_period_bt_prepaidCount");
    new Button({label: "充值",
        onClick: function(){
            // Do something:
            // dom.byId("result1").innerHTML += "Thank you! ";
        }
    }, "profile_period_bt_prepaid");
    new Button({label: "消费",
        onClick: function(){
            // Do something:
            // dom.byId("result1").innerHTML += "Thank you! ";
        }
    }, "profile_period_bt_consumption");

    var chart = new Chart("periodAnalyzeChart",{
        title: "ACU/PCU数据对比",
        titlePos: "top",
        titleGap: 25,
        titleFont: "normal normal normal 15pt Arial",
        titleFontColor: "orange"
    });

    chart.addPlot("default", {type: Lines});
    chart.addPlot("other", {type: Areas, hAxis: "other x", vAxis: "other y"});
    chart.addAxis("x");
    chart.addAxis("y", {vertical: true});
    chart.addAxis("other x", {leftBottom: false});
    chart.addAxis("other y", {vertical: true, leftBottom: false});
    chart.addSeries("Series 1", [1, 2, 2, 3, 4, 5, 5, 7]);
    chart.addSeries("Series 2", [1, 1, 4, 2, 1, 6, 4, 3],
        {plot: "other", stroke: {color:"blue"}, fill: "lightblue"}
    );
    chart.render();
}


});