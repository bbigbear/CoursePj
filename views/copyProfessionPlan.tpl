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
				<div class="col-sm-10">
					<div class="panel panel-primary">
						<div class="panel-heading">
							<h4 class="panel-title">复制专业培养计划</h4>
						</div>
					    <div class="panel-body">
						<form class="form-inline" role="form" id="searchPm">
					        <div class="form-group">				
								<label>年级</label>			
								<select class="form-control" name="year" id="year">
								<option>2015</option>
								<option>2016</option>
								<option>2017</option>
								<option>2018</option>
								</select>							
								<label>院系</label>
								<select class="form-control" name="faculty" id="faculty">
								{{range .m}}
								<option>{{.}}</option>
								{{end}}				
								</select>			
								<button type="button" class="btn btn-default" onclick="return QueryInput()">检索</button>				
							</div>
						</form>						
						</div>																	
					</div>					
				</div>				
      			<div class="col-sm-5">
					<div class="panel panel-primary">
					    <div class="panel-body">
						<div class="row">				
					    <div class="col-sm-8">							
							<form role="form">
							  <div class="form-group">
								<label for="name" id="f">院/系：{{.f}}</label><br>
							    <label for="name">已设置培养方案的专业列表：{{.len}}个</label>
							    <select multiple class="form-control" id="plan_list">
								{{range .s}}	
							      <option>{{.}}</option>							   
								{{end}}	
							    </select>
							  </div>
							</form>																			
						</div>
						<div class="col-sm-2" style="padding-top:50px">
							<button type="button" class="btn btn-primary" onclick="return CopyInput()">复制</button>
						</div>
						</div>						
						</div>
					</div>					
				</div>
				<div class="col-sm-5">
					<div class="panel panel-primary">
						<div class="panel-body">
						<div class="row">				
						    <div class="col-sm-8">							
								<form role="form">
								  <div class="form-group">
								    <label for="name">已制定开课专业列表：{{.slice_plan_len}}个</label>
								    <select multiple class="form-control" id="open_class_list">
									{{range .slice_plan}}	
								      <option>{{.}}</option>							   
									{{end}}	
								    </select>
								  </div>
								</form>																		
							</div>
							<div class="col-sm-2" style="padding-top:25px">
								<button type="button" class="btn btn-primary" onclick="return RemoveInput()">移除</button>
							</div>
						</div>						
					</div>
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
				if({{.f}}!=""){
					$("#faculty").val({{.f}})
				}													
				//alert("自动加载")			
			})
			
			function QueryInput(){
				var year=document.getElementById("year")
				var faculty=document.getElementById("faculty")
				if (faculty.value==""){
					alert("请输入学院")
				}else{
					window.location.href="/copyplan/profession/search?year="+year.value+"&faculty="+faculty.value
				}
			
			}
			function CopyInput(){
				//alert("点击复制按钮")
				//alert($("#plan_list").val())
				var year=document.getElementById("year")
				//var faculty=document.getElementById("f")
				var pmname=document.getElementById("plan_list")
				//alert(pmname.value)
				if($("#plan_list").val()!=null){
					$.ajax({  
					    url: "{{urlfor "CopyPlanController.PPCopy"}}",  
					    data: { 
							pmname: pmname.value,
							year: year.value,
							faculty: {{.f}}
						},    
					    type: "POST",
						async:false,
						error:function(data){
							alert("post error")
						},
					    success:function(data){  
					        if(data.status==0){
								alert("复制成功")
								window.location.href="/copyplan/profession/search?year="+year.value+"&faculty="+{{.f}}
							}else{
								alert("复制失败，已存在专业")
							}
					    }  
					});
				}else{
					alert("请选择专业再点击复制")
				}
				
			}
			function RemoveInput(){
				//alert("点击移除按钮")
				var pmname=document.getElementById("open_class_list")
				if($("#open_class_list").val()!=null){
					$.ajax({  
					    url: "{{urlfor "CopyPlanController.PPRemove"}}",  
					    data: { 
							pmname: pmname.value,
						},    
					    type: "POST",
						async:false,
						error:function(data){
							alert("post error")
						},
					    success:function(data){  
					        if(data.status==0){
								alert("移除成功")
								window.location.href="/copyplan/profession/search?year="+{{.y}}+"&faculty="+{{.f}}
							}else{
								alert("移除失败")
							}
					    }  
					});
				}else{
					alert("请选择专业再点击移除")
				}
			}
			
		</script>
	</body>
		
</html>
