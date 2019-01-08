/* --------layout start-------- */
function change_pwd(){
    var new_pwd = $('input[name=new_pwd]').val();
    var old_pwd = $('input[name=old_pwd]').val();
    var con_pwd = $('input[name=con_pwd]').val();
    $.ajax({
        url:'/admin/service/password_change',
        dataType:'json',
        type:'POST',
        cache:false,
        data:{new_pwd:new_pwd, old_pwd:old_pwd, con_pwd:con_pwd},
        success:function(data){
            if(data == null){
                swal("修改失败", "服务器错误", "error");
                return;
            }
            if (data.status != 0){
                swal("修改失败", data.msg, "error");
                return;
            }
            swal({title:"修改成功!",text: data.msg, type:"success" }, function(){
                window.location.href = document.referrer
            });
        }
    })
}

/* --------layout end-------- */


/*------login start------*/
function login_out(){
    $.ajax({
		url:"/admin/service/loginout",
		dataType:"json",
		type:"POST",
		cache:false,
		success:function(data){
			if (data.status == 0){
				window.location.href = "/admin/login";
			}
		},
	})
}
/*------login end------*/

/*------users start------*/
function user_add(){
    swal({ 
            title: "Add New User", 
            text: "Please input user name：",
            type: "input", 
            showCancelButton: true, 
            closeOnConfirm: false, 
            animation: "slide-from-top", 
            inputPlaceholder: "user name" 
        },
        function(inputValue){ 
            if (inputValue === false) return false; 
            
            if (inputValue === "") { 
                swal.showInputError("Please input user name！");
                return false 
            } 

            $.ajax({
                url:"/admin/service/user_add",
                dataType:'json',
                type:"POST",
                cache:false,
                data:{username:inputValue},
                success:function(data){
                    if(data == null){
                        swal("添加失败", "服务器错误", "error");
                        return;
                    }
                    if (data.status != 0){
                        swal("添加失败", data.msg, "error");
                        return;
                    }
                    swal({title:"添加成功!",text: "添加用户:【" + inputValue + "】成功!", type:"success" }, function(){
                        location.reload();
                    });
                }
            })
        }
    );
}

function user_edit(){
    var level = $('select[name=level]').val();
    var id = $('input[name=id]').val();
    $.ajax({
        url:'/admin/service/user_edit',
        dataType:'json',
        type:'POST',
        cache:false,
        data:{id:id, level:level},
        success:function(data){
            if(data == null){
                swal("修改失败", "服务器错误", "error");
                return;
            }
            if (data.status != 0){
                swal("修改失败", data.msg, "error");
                return;
            }
            swal({title:"修改成功!",text: data.msg, type:"success" }, function(){
                window.location.href = document.referrer
            });
        }
    })
}

function user_delete(id){
    swal({ 
        title: "确定删除吗？", 
        text: "你将无法恢复该用户！", 
        type: "warning",
        showCancelButton: true, 
        confirmButtonColor: "#DD6B55",
        confirmButtonText: "确定删除！", 
        closeOnConfirm: false
    },
    function(){
        $.ajax({
            url:'/admin/service/user_delete',
            dataType:'json',
            type:'POST',
            cache:false,
            data:{id:id},
            success:function(data){
                if(data == null){
                    swal("删除失败", "服务器错误", "error");
                    return;
                }
                if (data.status != 0){
                    swal("删除失败", data.msg, "error");
                    return;
                }
                swal({title:"删除成功!",text: data.msg, type:"success" }, function(){
                    location.reload();
                });
            }
        })
    });
}

function user_enable(id){
    swal({ 
        title: "确定启用该用户吗？", 
        text: "", 
        type: "warning",
        showCancelButton: true, 
        confirmButtonColor: "#DD6B55",
        confirmButtonText: "确定启用！", 
        closeOnConfirm: false
    },
    function(){
        $.ajax({
            url:'/admin/service/user_enable',
            dataType:'json',
            type:'POST',
            cache:false,
            data:{id:id},
            success:function(data){
                if(data == null){
                    swal("启用失败", "服务器错误", "error");
                    return;
                }
                if (data.status != 0){
                    swal("启用失败", data.msg, "error");
                    return;
                }
                swal({title:"启用成功!",text: data.msg, type:"success" }, function(){
                    location.reload();
                });
            }
        })
    });
}

function user_disable(id){
    swal({ 
        title: "确定禁用该用户吗？", 
        text: "", 
        type: "warning",
        showCancelButton: true, 
        confirmButtonColor: "#DD6B55",
        confirmButtonText: "确定禁用！", 
        closeOnConfirm: false
    },
    function(){
        $.ajax({
            url:'/admin/service/user_disable',
            dataType:'json',
            type:'POST',
            cache:false,
            data:{id:id},
            success:function(data){
                if(data == null){
                    swal("禁用失败", "服务器错误", "error");
                    return;
                }
                if (data.status != 0){
                    swal("禁用失败", data.msg, "error");
                    return;
                }
                swal({title:"禁用成功!",text: data.msg, type:"success" }, function(){
                    location.reload();
                });
            }
        })
    });
}
/*------users end------*/

/* -------app start------- */
function app_add(){
    var appname = $('input[name=appname]').val();
    var desc = $('input[name=desc]').val()
    $.ajax({
        url:'/admin/service/app_add',
        dataType:'json',
        type:'POST',
        cache:false,
        data:{appname:appname, desc:desc},
        success:function(data){
            if(data == null){
                swal("添加失败", "服务器错误", "error");
                return;
            }
            if (data.status != 0){
                swal("添加失败", data.msg, "error");
                return;
            }
            swal({title:"添加成功!",text: "添加应用:【" + appname + "】成功!", type:"success" }, function(){
                window.location.href = document.referrer
            });
        }
    })
}

function app_delete(id){
    swal({ 
        title: "确定删除吗？", 
        text: "你将无法恢复该应用！", 
        type: "warning",
        showCancelButton: true, 
        confirmButtonColor: "#DD6B55",
        confirmButtonText: "确定删除！", 
        closeOnConfirm: false
    },
    function(){
        $.ajax({
            url:'/admin/service/app_delete',
            dataType:'json',
            type:'POST',
            cache:false,
            data:{id:id},
            success:function(data){
                if(data == null){
                    swal("删除失败", "服务器错误", "error");
                    return;
                }
                if (data.status != 0){
                    swal("删除失败", data.msg, "error");
                    return;
                }
                swal({title:"删除成功!",text: data.msg, type:"success" }, function(){
                    location.reload();
                });
            }
        })
    });
}

function app_apply(id){
    swal({ 
        title: "确定申请激活该应用吗？", 
        text: "", 
        type: "warning",
        showCancelButton: true, 
        confirmButtonColor: "#DD6B55",
        confirmButtonText: "确定申请！", 
        closeOnConfirm: false
    },
    function(){
        $.ajax({
            url:'/admin/service/user_apply',
            dataType:'json',
            type:'POST',
            cache:false,
            data:{id:id},
            success:function(data){
                if(data == null){
                    swal("申请失败", "服务器错误", "error");
                    return;
                }
                if (data.status != 0){
                    swal("申请失败", data.msg, "error");
                    return;
                }
                swal({title:"申请成功!",text: data.msg, type:"success" }, function(){
                    location.reload();
                });
            }
        })
    });
}

function app_pass(id){
    swal({ 
        title: "确定审核通过吗？", 
        text: "", 
        type: "warning",
        showCancelButton: true, 
        confirmButtonColor: "#DD6B55",
        confirmButtonText: "确定通过！", 
        closeOnConfirm: false
    },
    function(){
        $.ajax({
            url:'/admin/service/app_pass',
            dataType:'json',
            type:'POST',
            cache:false,
            data:{id:id},
            success:function(data){
                if(data == null){
                    swal("审核失败", "服务器错误", "error");
                    return;
                }
                if (data.status != 0){
                    swal("审核失败", data.msg, "error");
                    return;
                }
                swal({title:"审核成功!",text: data.msg, type:"success" }, function(){
                    location.reload();
                });
            }
        })
    });
}

function app_unpass(id){
    swal({ 
        title: "确定审核不通过吗？", 
        text: "", 
        type: "warning",
        showCancelButton: true, 
        confirmButtonColor: "#DD6B55",
        confirmButtonText: "确定不通过！", 
        closeOnConfirm: false
    },
    function(){
        $.ajax({
            url:'/admin/service/app_unpass',
            dataType:'json',
            type:'POST',
            cache:false,
            data:{id:id},
            success:function(data){
                if(data == null){
                    swal("审核失败", "服务器错误", "error");
                    return;
                }
                if (data.status != 0){
                    swal("审核失败", data.msg, "error");
                    return;
                }
                swal({title:"审核成功!",text: data.msg, type:"success" }, function(){
                    location.reload();
                });
            }
        })
    });
}
/* -------app end------- */