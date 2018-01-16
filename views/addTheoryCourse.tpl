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
							<h3 class="panel-title">录入理论课程信息</h3>
						</div>
					    <div class="panel-body">
						<form class="form-horizontal" method="POST" id="addCourse">
					        <div class="form-group">				
								<label class="col-sm-2 control-label">课程代码</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Cid" id="cid">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">开课单位</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Cunit">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">课程名称</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Cname">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">课程类别</label>			
								<div class="col-sm-5">
								<select class="form-control" name="Ccg1">
								<option>计算机</option>
								<option>音乐</option>
								<option>外语</option>
								<option>大学语文</option>
								</select>
								</div>
								<div class="col-sm-5">
								<select class="form-control" name="Ccg2">
								<option>选修</option>
								<option>必修</option>
								</select>							
								</div>
								
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">英文名称</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Cname_en">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">状态</label>			
								<div class="col-sm-10">
								<select class="form-control" name="Status">
								<option>可用</option>
								<option>停用</option>
								</select>
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">学分</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Credit">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">授课学时</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Tteach">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">实验学时</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Texperiment">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">上机学时</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Tcomputer">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">其他学时</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Tother">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">总学时</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Ttotal">
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
								<select class="form-control" name="Year">
								<option>2015</option>
								<option>2016</option>
								<option>2017</option>
								<option>2018</option>
								</select>
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
				var cid=document.getElementById("cid")
				if(cid.value.length==0){
					alert("课程代码不能为空")
					return false
				}
				
				$.ajax({
					type:"POST",
					url:"{{urlfor "HomeController.TheoryCourseAddAction"}}",
					data:$("#addCourse").serialize(),
					async:false,
					error:function(request){
						alert("post error")
						
					},
					success:function(data){
						if(data.status==0){
							alert("新增成功")
							window.location.href="/home/add"
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
