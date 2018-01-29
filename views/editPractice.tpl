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
							<h3 class="panel-title">编辑实践环节信息</h3>
						</div>
					    <div class="panel-body">
						{{range .m}}
						<form class="form-horizontal" role="form" id="editPractice">
							<input type="hidden" name="Id" value="{{.Id}}">
					        <div class="form-group">				
								<label class="col-sm-2 control-label">环节代码</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Pid" value="{{.Pid}}" id="pid">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">开课单位</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Punit" value="{{.Punit}}">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">环节名称</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Pname" value="{{.Pname}}">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">环节类别</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Pcg1" value="{{.Pcg1}}">					
								</div>	
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">英文名称</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Pname_en" value="{{.Pname_en}}">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">状态</label>			
								<div class="col-sm-10">
								<select class="form-control" name="Status" id="Status" value="{{.Status}}">
								<option>可用</option>
								<option>停用</option>
								</select>
<!--								<input class="form-control" name="Status" value="{{.Status}}">-->
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">学分</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Credit" value="{{.Credit}}" type="number">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">学时</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Tclass" value="{{.Tclass}}" type="number">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">周数</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Nw" value="{{.Nw}}" type="number">
								</div>
							</div>			
							<div class="form-group">				
								<label class="col-sm-2 control-label">教学大纲</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Syllabus" value="{{.Syllabus}}">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">年级</label>			
								<div class="col-sm-10">
								<select class="form-control" name="Year" id="year" value="{{.Year}}">
								<option>2015</option>
								<option>2016</option>
								<option>2017</option>
								<option>2018</option>
								</select>
<!--								<input class="form-control" name="Year" value="{{.Year}}" id="year">-->
								</div>
							</div>
							<div class="form-group">							
								<div class="col-sm-2 pull-right">
									<button type="button" class="btn btn-primary" onclick="return UpdataInput()">修改</button>
									<button type="button" class="btn btn-primary" onclick="return DeleteInput()">删除</button>
								</div>
							</div>
						</form>
						{{end}}	
					</div>
				</div>
			</div>			
		</div>
		<script type="text/javascript">
			//自动加载
			$(function(){
				if({{.y}}!=""){
					$("#year").val({{.y}})
				}
				if({{.s}}!=""){
					$("#Status").val({{.s}})
				}													
				//alert("自动加载")			
			})
			
			function UpdataInput(){
				$.ajax({
					type:"POST",
					url:"{{urlfor "PracticeController.PracticeUpdata"}}",
					data:$("#editPractice").serialize(),
					async:false,
					error:function(request){
						alert("post error")				
					},
					success:function(data){
						if(data.status==0){
							alert("更新成功")
							window.location.href="/practice"
						}else{
							alert("更新失败")
						}
						
					}
					
				});
				return true	
			}
			function DeleteInput(){
				var pid=document.getElementById("pid")
				var year=document.getElementById("year")
				$.ajax({
					type:"POST",
					url:"{{urlfor "PracticeController.PracticeDelete"}}",
					data:{pid:pid.value,year:year.value},
					async:false,
					error:function(request){
						alert("post error")				
					},
					success:function(data){
						if(data.status==0){
							alert("删除成功")
							window.location.href="/practice"
						}else{
							alert("删除失败")
						}
						
					}
					
				});
				return true	
			}	
		</script>
	</body>
		
</html>
