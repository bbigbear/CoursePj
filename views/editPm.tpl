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
						<form class="form-horizontal" role="form" id="editPm">
							<input type="hidden" name="Id" value="{{.Id}}">
					        <div class="form-group">				
								<label class="col-sm-2 control-label">专业代码</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Pmid" id="pmid" value="{{.Pmid}}">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">专业名称</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Pmname" value="{{.Pmname}}">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">院系</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Faculty" value="{{.Faculty}}">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">培养层次</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Train_level" value="{{.Train_level}}">
								</div>
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">允许辅修</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Isminor" value="{{.Isminor}}">
								</div>					
							</div>
							<div class="form-group">				
								<label class="col-sm-2 control-label">英文名称</label>			
								<div class="col-sm-10">
								<input class="form-control" name="Pmname_en" value="{{.Pmname_en}}">
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
					url:"/pm/updata",
					data:$("#editPm").serialize(),
					async:false,
					error:function(request){
						alert("post error")				
					},
					success:function(data){
						if(data.status==200){
							alert(data.message)
							window.location.href="/pm"
						}else{
							alert(data.message)
						}
						
					}
					
				});
				return true	
			}
			function DeleteInput(){
				var pmid=document.getElementById("pmid")
				var year=document.getElementById("year")
				$.ajax({
					type:"POST",
					//contentType:"application/json;charset=utf-8",
					url:"/pm/delete",
					data:{pmid:pmid.value,year:year.value},
					//data:JSON.stringify({'pmid':pmid.value,'year':parseInt(year.value)}),
					async:false,
					error:function(request){
						alert("post error")				
					},
					success:function(data){
						if(data.status==200){
							alert(data.message)
							window.location.href="/pm"
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
