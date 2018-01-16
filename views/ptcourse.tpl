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
							<h4 class="panel-title">查询专业理论课程</h4>
						</div>
					    <div class="panel-body">
						<form class="form-inline" role="form" id="searchCourse">					        
							<div class="form-group">						
								<label>所属的年级</label>
								<input class="form-control" name="s_Year" id="s_Year">
								<label>院系</label>
							  	<input class="form-control" name="s_Faculty" id="s_Faculty">			
								<button type="button" class="btn btn-primary" onclick="return QueryInput()">检索</button>
							</div>
						</form>
						</div>
					</div>
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
						<button type="button" class="btn btn-primary" onclick="return ToSetCourse()">设置专业课程</button>
						<button type="button" class="btn btn-primary" onclick="return QueryInput()" style="margin-top:10px">查看已设置课程</button>
					</div>
					</div>
					<div class="table-responsive">
						<table class="table table-bordered">
							<caption><h4 class="panel-title" >理论课程信息</h4></caption>
							<thead>
								<tr>
									<th>选择</th>
									<th>课程代码</th>
									<th>开课单位</th>
									<th>课程名称</th>
									<th>课程类别1</th>
									<th>课程类别2</th>
									<th>英文名称</th>
									<th>状态</th>
									<th>学分</th>
									<th>授课学时</th>
									<th>实验学时</th>
									<th>上机学时</th>
									<th>其他学时</th>
									<th>总学时</th>
									<th>年级</th>
									<th>教学大纲</th>
								</tr>
								{{range .m}}
								<tr>
									<th><input type="checkbox" value="{{.Cid}}" name="Cid"></th>
									<th>{{.Cid}}</th>
									<th>{{.Cunit}}</th>
									<th>{{.Cname}}</th>
									<th>{{.Ccg1}}</th>
									<th>{{.Ccg2}}</th>
									<th>{{.Cname_en}}</th>
									<th>{{.Status}}</th>
									<th>{{.Credit}}</th>
									<th>{{.Tteach}}</th>
									<th>{{.Texperiment}}</th>
									<th>{{.Tcomputer}}</th>
									<th>{{.Tother}}</th>
									<th>{{.Ttotal}}</th>
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
			
			function QueryInput(){				
				var s_Status=document.getElementById("s_Status")
				var s_Year=document.getElementById("s_Year")
				window.location.href="/home/search?s_Year="+s_Year.value+"&s_Status="+s_Status.value
			}
			function ToSetCourse(){
				var checked_array=[];
				var selectedValues=[];
				var data="";
				var Pm_data="";
			 	$("[name='Cid']:checkbox:checked").each(function(){				 
					checked_array.push($(this).val()) 	
					data=data+$(this).val()+',';			
				});
				$("#Pmid :selected").each(function(){
				     selectedValues.push($(this).val());
					 Pm_data=Pm_data+$(this).val()+',';
				 });
				alert(Pm_data)
				alert(data)
				
				$.ajax({  
				    url: "{{urlfor "PTCourseController.Setcourse"}}",  
				    data: { cid: data,pmid: Pm_data},    
				    type: "POST",   
				    success: function () {  
				        
				    }  
				});			
			}		
		</script>
	</body>
		
</html>
