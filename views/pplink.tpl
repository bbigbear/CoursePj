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
							<h4 class="panel-title">设置专业实践环节</h4>
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
								{{range .f}}
								<option>{{.}}</option>
								{{end}}				
								</select>			
								<button type="button" class="btn btn-default" onclick="return QueryInput()">检索</button>				
							</div>
						</form>						
						</div>																	
					</div>					
				</div>
				<div class="col-sm-10">
					<div class="panel panel-primary">						
					    <div class="panel-body">
						<!--<form class="form-inline" role="form" id="searchCourse">					        
							<div class="form-group">						
								<label>所属的年级</label>
								<input class="form-control" name="s_Year" id="s_Year">
								<label>院系</label>
							  	<input class="form-control" name="s_Faculty" id="s_Faculty">			
								<button type="button" class="btn btn-primary" onclick="return QueryInput()">检索</button>
							</div>
						</form>-->
						
						<div class="row">				
					    <div class="col-sm-4">							
							<form role="form">
							  <div class="form-group">
							    <label for="name">专业列表</label>
							    <select multiple class="form-control" id="Pmid">
								{{range .m1}}	
							      <option>{{.Pmname}}[{{.Pmid}}]</option>							   
								{{end}}	
							    </select>
							  </div>
							</form>																			
						</div>
						<div class="col-sm-2" style="padding-top:25px">
							<button type="button" class="btn btn-primary" onclick="return ToSetCourse()">设置实践环节</button>
							<button type="button" class="btn btn-primary" onclick="return EditInput()" style="margin-top:10px">查看已设置环节</button>
						</div>
						</div>
						
						</div>
					</div>
					
					<div class="table-responsive">
						<table class="table table-bordered">
							<caption><h4 class="panel-title">实践环节信息</h4></caption>
							<thead>
								<tr>
									<th>选择</th>
									<th>环节代码</th>
									<th>开课单位</th>
									<th>环节名称</th>
									<th>环节类别1</th>
									<th>英文名称</th>
									<th>状态</th>
									<th>学分</th>
									<th>学时</th>
									<th>周数</th>
									<th>年级</th>
									<th>教学大纲</th>
								</tr>
								{{range .m}}
								<tr>
									<th><input type="checkbox" value="{{.Pid}}" name="Pid" id="Pid"></th>
									<th>{{.Pid}}</th>
									<th>{{.Punit}}</th>
									<th>{{.Pname}}</th>
									<th>{{.Pcg1}}</th>
									<th>{{.Pname_en}}</th>
									<th>{{.Status}}</th>
									<th>{{.Credit}}</th>
									<th>{{.Tclass}}</th>
									<th>{{.Nw}}</th>
									<th>{{.Year}}</th>
									<th>{{.Syllabus}}</th>
								</tr>							
								{{end}}					
							</thead>
							<tbody>							
							</tbody>
						</table>
					</div> 
				</div>
      		</div>
    	</div>
		<script type="text/javascript">
			
			function EditInput(){				
				var pmid=document.getElementById("Pmid")
				if ($("#Pmid").val()!=null){
					window.location.href="/pplink/edit?pmid="+pmid.value
				}else{
					alert("请选择专业再查看")
				}
			}
			function ToSetCourse(){
				var checked_array=[];
				var selectedValues=[];
				var data="";
				var Pm_data="";
			 	$("[name='Pid']:checkbox:checked").each(function(){				 
					checked_array.push($(this).val()) 	
					data=data+$(this).val()+',';			
				});
				$("#Pmid :selected").each(function(){
				     selectedValues.push($(this).val());
					 Pm_data=Pm_data+$(this).val()+',';
				 });
				//alert(Pm_data)
				//alert(data)
				if($("[name='Pid']:checked").length>0&&$("#Pmid").val()!=null){
					$.ajax({  
					    url: "{{urlfor "PPLinkController.Setcourse"}}",  
					    data: { pid: data,pmid: Pm_data},    
					    type: "POST",
						async:false,
						error:function(data){
							alert("post error")
						},
					    success:function(data){  
					        if(data.status==0){
								alert("设置成功")
							}else{
								alert("设置失败，已存在相关实践环节")
							}
					    }  
					});
				}else{
					alert("不能为空")
				}			
							
			}		
		</script>
	</body>
		
</html>
