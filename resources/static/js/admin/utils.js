//统一表单
function getTableOption(elem, url, cols, where, method) {
    if (where === undefined) {
        where = {};
    }
    if (method === undefined) {
        method = 'get';
    }
    var tableOption = {
        // skin: 'line', //行边框风格
        loading: true,
        where: where,
        // even: true, //开启隔行背景
        // size: 'sm', //小尺寸的表格
        size: 'lg', //小尺寸的表格
        elem: elem,
//            height: 315,
        url: url, //数据接口
        cellMinWidth: 80, //全局定义常规单元格的最小宽度，layui 2.2.1 新增
        page: true,//开启分页
        method: method,
        // limit: 1,
        // limits: [10,20,30,40,50,60,70,80,90],
        request: {
            pageName: 'page',
            limitName: 'limit',

        },
        parseData: function (res) {
            console.log("parseData======================")
            console.log(res)
        },
        response: {
            statusName: 'error_code',
            statusCode: 0,
            countName: 'count',
            dataName: 'list',
            msgName: 'error_msg'
        },
        errorHandler: function (e, m) {
            if (e.status === 401) {
                if (window != top) {
                    top.location.href = "/admin/login";
                } else {
                    window.location.href = "/admin/login";
                }

            }
        },
        cols: [cols],
        done: function (res, curr, count) {
            console.log(res)
            //数据渲染后回调
            //如果是异步请求数据方式，res即为你接口返回的信息。
            //如果是直接赋值的方式，res即为：{data: [], count: 99} data为当前页数据、count为数据总长度
            // console.log(res);
            //得到当前页码
            // console.log(curr);
            //得到数据总量
            // console.log(count);
        }
    };
    return tableOption;
}

function tableInit(table, elem, url, cols, where) {
    table.render(getTableOption(elem, url, cols, where));
}


function ajaxRequest(url, data, type, successHandler = undefined, errorHandler = undefined, async = false) {
    console.log(url)
    console.log(data)
    console.log(type)
    console.log(successHandler)
    console.log(errorHandler)
    // $.ajax({
    //     type: type,
    //     url: url,
    //     async: async,
    //     data: data,
    //     success: function (res) {
    //         successHandler(res);
    //     },
    //     error: function (res) {
    //         errorHandler(res)
    //     }
    // });
}