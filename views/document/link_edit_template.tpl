<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <title>编辑文档 </title>

    <!-- Bootstrap -->
    <link href="{{cdncss "/static/bootstrap/css/bootstrap.min.css"}}" rel="stylesheet" type="text/css">
    <link href="{{cdncss "/static/font-awesome/css/font-awesome.min.css"}}" rel="stylesheet" type="text/css">

    <link href="{{cdncss "/static/css/main.css"}}" rel="stylesheet">

    <link rel="stylesheet" href="/static/zTree/css/zTreeStyle/zTreeStyle.css" type="text/css">

    <script src="{{cdnjs "/static/jquery/1.12.4/jquery.min.js"}}" type="text/javascript"></script>
    <script type="text/javascript" src="/static/zTree/js/jquery.ztree.core.js"></script>
    <script type="text/javascript" src="/static/zTree/js/jquery.ztree.excheck.js"></script>
    <script type="text/javascript" src="/static/zTree/js/jquery.ztree.exedit.js"></script>
    <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
    <script src="/static/html5shiv/3.7.3/html5shiv.min.js"></script>
    <script src="/static/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->
</head>
<body>
<div class="m-manual manual-editor">
    <div class="manual-editor-container" id="manualEditorContainer" style="top: 0;">
        <form role="form" method="post" name="editLinkForm" id="editLinkForm">
            <input type="hidden" name="link_docs" id="link_docs" value="{{.LinkDocLinks}}">
            <div class="form-group">
                <a href="{{urlfor "BookController.Index"}}" data-toggle="tooltip" data-title="返回"><i class="fa fa-chevron-left" aria-hidden="true">返回</i></a>
                <button type="submit" class="btn btn-success" data-loading-text="保存中...">保存修改</button>
                <span id="form-error-message" class="error-message"></span>
            </div>
            <div class="form-group">
                <div class="zTreeDemoBackground" style="margin-left: 20px;margin-right: 0px">
                    <ul id="treeDemo" class="ztree" style="height:320px; width: auto; margin-bottom: 10px;">
                    </ul>
                </div>
                </ul>
            </div>
        </form>
    </div>
</div>
<script src="{{cdnjs "/static/bootstrap/js/bootstrap.min.js"}}" type="text/javascript"></script>
<script src="{{cdnjs "/static/vuejs/vue.min.js"}}" type="text/javascript"></script>
<script src="{{cdnjs "/static/js/jquery.form.js"}}" type="text/javascript"></script>
<script src="/static/js/main.js" type="text/javascript"></script>

<script type="text/javascript">
    <!--
    var setting = {
        check: {
            enable: true
        },
        data: {
            simpleData: {
                enable: true
            }
        },
        callback: {
            onCheck: onCheck
        }
    };
    var zNodes = JSON.parse({{.LinkDocResult}});
    var selectNodes = "{{.LinkDocLinks}}";

    $(document).ready(function(){
        $.fn.zTree.init($("#treeDemo"), setting, zNodes);
        setCheck();
        $("#py").bind("change", setCheck);
        $("#sy").bind("change", setCheck);
        $("#pn").bind("change", setCheck);
        $("#sn").bind("change", setCheck);
    });


    function setCheck() {
        type = { "Y" : "ps", "N" : "ps" };
        setting.check.chkboxType = type;
    }
    function onCheck(e, treeId, treeNode) {
        resetDocs();
    }
    function resetDocs() {
        var treeObj = $.fn.zTree.getZTreeObj("treeDemo");
        var nodes = treeObj.getCheckedNodes(true);
        selectNodes = '';
        $.each(nodes,function(k,v){
            selectNodes += v['id']+","
        });
        $("#link_docs").val(selectNodes);
        showSuccess($("#link_docs").val());
    }
    $("#editLinkForm").ajaxForm({
        beforeSubmit : function () {
            resetDocs();
        },
        success : function (res) {
            //showSuccess(JSON.stringify(res));
            if(res.errcode === 7000){
                showError('失败：'+res.message);
            }else{
                showSuccess('保存成功');
            }
        }
    }) ;

    //-->
</script>

</body>
</html>