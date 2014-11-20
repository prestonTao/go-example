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
    // localStore();
    serverStore();
});


function localStore(){
    //定义客户端store
    var store = new Store({  
        data: [  
            {id: 1, title: 'Hey There', artist: 'Bette Midler'},  
            {id: 2, title: 'Love or Confusion', artist: 'Jimi Hendrix'},  
            {id: 3, title: 'Sugar Street', artist: 'Andy Narell'}  
        ]  
    });

    //定义列结构
    var columns = [  
        {field: 'id', name: 'Identity'},  
        {field: 'title', name: 'Title'},  
        {field: 'artist', name: 'Artist'}  
    ]; 

    var grid = new Grid({  
        cacheClass: Sync,  
        store: store,  
        structure: columns,
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
    }, 'gridNode');
    grid.startup(); 
}


function serverStore(){


    var store = new JsonRestStore({target: "/gridxdemo/" });

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
    }, 'asyncGrid');
    grid.startup(); 
}



});