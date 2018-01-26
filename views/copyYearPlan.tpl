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
							<h4 class="panel-title">复制年级教学计划</h4>
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
								<label for="name" id="y">年级：{{.y}}</label>
								<label for="name" id="f">院/系：{{.f}}</label><br>
							    <label for="name">已制定开课专业列表：{{.l}}个</label>
							    <select multiple class="form-control" id="left_list">
								{{range .maps}}	
							      <option>{{.Plname}}</option>							   
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
									<label class="col-sm-4" style="padding-top:10px;padding-left:3px">选择要复制年级</label>											
									<select class="form-control" name="year" id="year_right" style="width:80px">
										<option>2015</option>
										<option>2016</option>
										<option>2017</option>
										<option>2018</option>
									</select>								
									<br>
								    <label for="name">已制定开课专业列表：{{.l1}}个</label>
								    <select multiple class="form-control" id="right_list">
									{{range .maps_right}}	
								      <option>{{.Plname}}</option>							   
									{{end}}	
								    </select>
								  </div>
								</form>																		
							</div>
							<div class="col-sm-2" style="padding-left:0px">
								<button type="button" class="btn btn-primary" onclick="return SearchInput()">检索选择年级已制定专业</button>
							</div>
							<div class="col-sm-2" style="padding-top:80px">
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
				if({{.year_right}}!="")	{
					$("#year_right").val({{.year_right}})
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
					window.location.href="/copyplan/year/search?year="+year.value+"&faculty="+faculty.value
				}
			}
			function SearchInput(){				
				var year_right=document.getElementById("year_right")
				//$("#year_right").val(year_right.value)
				window.location.href="/copyplan/year/search?year="+year.value+"&faculty="+faculty.value+"&year_right="+year_right.value
											
			}
			function CopyInput(){
				//alert("点击复制按钮")
				//alert($("#plan_list").val())
				var year=document.getElementById("year")
				//var faculty=document.getElementById("f")
				var year_right=document.getElementById("year_right")
				var plname=document.getElementById("left_list")
				//alert(pmname.value)
				if(year.value!=year_right.value){
					if($("#left_list").val()!=null){
					$.ajax({  
					    url: "{{urlfor "CopyPlanController.GYCopy"}}",  
					    data: { 
							plname: plname.value,
							year: year_right.value,
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
								window.location.href="/copyplan/year/search?year="+year.value+"&faculty="+{{.f}}+"&year_right="+year_right.value
							}else{
								alert("复制失败，已存在专业")
							}
					    }  
					});
					}else{
						alert("请选择专业再点击复制")
					}
				}else{
					alert("请在右边选择不同的年级进行复制")
				}
				
				
			}
			function RemoveInput(){
				//alert("点击移除按钮")
				var pmname=document.getElementById("right_list")
				var year_right=document.getElementById("year_right")
				if($("#right_list").val()!=null){
					$.ajax({  
					    url: "{{urlfor "CopyPlanController.GYRemove"}}",  
					    data: { 
							pmname: pmname.value,
							year: year_right.value,
						},    
					    type: "POST",
						async:false,
						error:function(data){
							alert("post error")
						},
					    success:function(data){  
					        if(data.status==0){
								alert("移除成功")
								window.location.href="/copyplan/year/search?year="+year.value+"&faculty="+{{.f}}
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
