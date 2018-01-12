<!DOCTYPE html>
<html>
	<head>
		{{template "header"}}
		<script src="http://cdn.static.runoob.com/libs/jquery/2.1.1/jquery.min.js"></script>
		<script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
	</head>
	<body>
		<div class="container-fluid" style="padding-top:50px">
      		<div class="row">
        		{{template "LeftBar"}}
				<div class="col-sm-9">
					<div class="panel panel-primary">
						<div class="panel-heading">
							<h3 class="panel-title">录入专业信息</h3>
						</div>
					    <div class="panel-body">
						<form class="form-horizontal" method="POST" id="addPm">
					        <div class="form-group">				
								<label class="col-sm-2 control-label">专业代码</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Pmid" id="pmid">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">专业名称</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Pmname">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">院系</label>			
								<div class="col-sm-10">
								<input class="form-control" name="faculty">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">培养层次</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Train_level">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">允许辅修</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Isminor">
								</div>					
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">英文名称</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Pmname_en">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">状态</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Status">
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
				var cid=document.getElementById("pmid")
				if(pmid.value.length==0){
					alert("实践环节不能为空")
					return false
				}
				
				$.ajax({
					type:"POST",
					url:"{{urlfor "PmController.PmAddAction"}}",
					data:$("#addPm").serialize(),
					async:false,
					error:function(request){
						alert("post error")
						
					},
					success:function(data){
						if(data.status==0){
							alert("新增成功")
							window.location.href="/pm/add"
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
