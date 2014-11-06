require([
    "dojo/ready",
    'dojo/store/Memory',  
    'gridx/Grid',  
    'gridx/core/model/cache/Sync',
    'gridx/core/model/cache/Async',
    "dojo/store/JsonRest",
    'gridx/allModules',
    'gridx/modules/SingleSort'
], function(ready, Store, Grid, Sync, Async, JsonRestStore, modules, SingleSort){

ready(function(){
    serverStore();
});




function serverStore(){


    var store = new JsonRestStore({target: "/roleLogingridx/" });

    //定义列结构
    var columns = [  
        {field: 'id', name: 'Identity'},  
        {field: 'title', name: 'Title'},  
        {field: 'artist', name: 'Artist'}  
    ]; 

    var grid = new Grid({  
        cacheClass: Async,
        store: store,
        structure: columns,
        pageSize:10, //每页显示多少条数据
        modules:[
                SingleSort,
                modules.VirtualVScroller,
                //modules.ColumnResizer,
                modules.CellWidget,
                modules.RowHeader,  //添加表格最前面的列
                modules.IndirectSelect,  //添加checkBox
                modules.ExtendedSelectRow,
                //TouchVScroller  //另一种滑动方式
                "gridx/modules/Pagination",
                "gridx/modules/pagination/PaginationBar"
            ]
    }, 'roleLoginTable');
    grid.startup(); 
}



});