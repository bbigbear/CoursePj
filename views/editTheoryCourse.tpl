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
							<h3 class="panel-title">编辑理论课程信息</h3>
						</div>
					    <div class="panel-body">
						{{range .m}}
						<form class="form-horizontal" role="form" id="editCourse">
							<input type="hidden" name="Id" value="{{.Id}}">
					        <div class="form-group">				
								<label class="col-sm-2 control-label">课程代码</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Cid" value="{{.Cid}}" id="cid">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">开课单位</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Cunit" value="{{.Cunit}}">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">课程名称</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Cname" value="{{.Cname}}">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">课程类别</label>			
								<div class="col-sm-5">
								<input class="form-control" name="Ccg1" value="{{.Ccg1}}">
								
								</div>
								<div class="col-sm-5">
								<input class="form-control" name="Ccg2" value="{{.Ccg2}}">								
								</div>
								
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">英文名称</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Cname_en" value="{{.Cname_en}}">
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
								<label class="col-sm-2 control-label">授课学时</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Tteach" value="{{.Tteach}}" type="number">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">实验学时</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Texperiment" value="{{.Texperiment}}" type="number">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">上机学时</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Tcomputer" value="{{.Tcomputer}}" type="number">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">其他学时</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Tother" value="{{.Tother}}" type="number">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">总学时</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Ttotal" value="{{.Ttotal}}" type="number">
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
					url:"/home/updata",
					data:$("#editCourse").serialize(),
					async:false,
					error:function(request){
						alert("post error")				
					},
					success:function(data){
						if(data.status==200){
							alert(data.message)
							window.location.href="/home"
						}else{
							alert(data.message)
						}
						
					}
					
				});
				return true	
			}
			function DeleteInput(){
				var cid=document.getElementById("cid")
				var year=document.getElementById("year")
				$.ajax({
					type:"POST",
					//contentType:"application/json;charset=utf-8",
					url:"/home/delete",
					data:{cid:cid.value,year:year.value},
					//data:JSON.stringify({'cid':cid.value,'year':parseInt(year.value)}),
					async:false,					
					error:function(request){
						alert("post error")				
					},
					success:function(data){
						if(data.status==200){
							alert(data.message)
							window.location.href="/home"
						}else{
							alert(data.message)
						}
						
					}
					
				});
				return true	
			}	
		</script>
	</body>
		
</html>
