<!DOCTYPE html>
<html>
	<head>
		{{template "header"}}
		<script src="http://cdn.static.runoob.com/libs/jquery/2.1.1/jquery.min.js"></script>
		<script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
	</head>
	<body>
		{{template "TopBar"}}
		<div class="container-fluid" style="padding-top:50px">
      		<div class="row">
        		{{template "LeftBar"}}
				<div class="col-sm-9">
					<div class="panel panel-primary">
						<div class="panel-heading">
							<h3 class="panel-title">录入实践环节信息</h3>
						</div>
					    <div class="panel-body">
						<form class="form-horizontal" method="POST" id="addPractice">
					        <div class="form-group">				
								<label class="col-sm-2 control-label">环节代码</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Pid" id="pid">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">开课单位</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Punit">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">环节名称</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Pname">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">环节类别</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Pcg1">
								</div>					
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">英文名称</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Pname_en">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">状态</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Status">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">学分</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Credit">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">学时</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Tclass">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">周数</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Nw">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">教学大纲</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Syllabus">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">年级</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Year">
								</div>
							</div>
							<div class="form-group">							
								<div class="col-sm-1 pull-right">
									<button type="button" class="btn btn-primary" onclick="return AddInput()">新建</button>
								</div>
							</div>
						</form>
					</div>
				</div>
			</div>			
		</div>
		<script type="text/javascript">
			function AddInput(){
				var cid=document.getElementById("pid")
				if(pid.value.length==0){
					alert("环节代码不能为空")
					return false
				}
				
				$.ajax({
					type:"POST",
					url:"{{urlfor "PracticeController.PracticeAddAction"}}",
					data:$("#addPractice").serialize(),
					async:false,
					error:function(request){
						alert("post error")
						
					},
					success:function(data){
						if(data.status==0){
							alert("新增成功")
							window.location.href="/practice/add"
						}else{
							alert("新增失败，已有重复的课程代码")
						}
						
					}
					
				});
				
				return true	
			}	
		</script>
	</body>
		
</html>
