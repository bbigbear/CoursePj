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
							<h4 class="panel-title">查询理论课程信息</h4>
						</div>
					    <div class="panel-body">
						<form class="form-inline" role="form" id="searchCourse">
					        <div class="form-group">				
								<label>开课单位</label>			
								<input class="form-control" name="s_Cunit" id="s_Cunit">
								<label>课程名称</label>			
								<input class="form-control" name="s_Cname" id="s_Cname">								
							</div>
							<div class="form-group">				
								<label>课程类别1</label>
								<input class="form-control" name="s_Ccg1" id="s_Ccg1">							
								<label>课程类别2</label>
								<input class="form-control" name="s_Ccg2" id="s_Ccg2">
							</div>
							<br>
							<div class="form-group">				
								<label>是否停用</label>
								<select class="form-control" name="s_Status" id="s_Status">
								<option>可用</option>
								<option>停用</option>
								</select>
<!--								<input class="form-control" name="s_Status" id="s_Status">-->
								<label>所属的年级</label>
								<select class="form-control" name="s_Year" id="s_Year">
								<option>2015</option>
								<option>2016</option>
								<option>2017</option>
								<option>2018</option>
								</select>
<!--								<input class="form-control" name="s_Year" id="s_Year">				-->
								<button type="button" class="btn btn-primary" onclick="return QueryInput()">检索</button>
							</div>
						</form>
						<div class="col-sm-1 pull-right">					
							<button type="button" class="btn btn-primary" onclick="return AddInput()">新增</button>																	 	
					    </div>
						</div>
					</div>
					<div class="table-responsive">
						<table class="table table-bordered">
							<caption><h4 class="panel-title">理论课程信息</h4></caption>
							<thead>
								<tr>
									<th>操作</th>
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
									<th><a href="/home/edit?cid={{.Cid}}&year={{.Year}}">编辑</a></th>
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
			//自动加载
			$(function(){
				if({{.y}}!=""){
					$("#s_Year").val({{.y}})
				}
				if({{.s}}!=""){
					$("#s_Status").val({{.s}})
				}													
				//alert("自动加载")			
			})
			
			function QueryInput(){
				var s_Cunit=document.getElementById("s_Cunit")
				var s_Cname=document.getElementById("s_Cname")
				var s_Ccg1=document.getElementById("s_Ccg1")
				var s_Ccg2=document.getElementById("s_Ccg2")
				var s_Status=document.getElementById("s_Status")
				var s_Year=document.getElementById("s_Year")
				window.location.href="/home/search?s_Cunit="+s_Cunit.value+"&s_Cname="+s_Cname.value+"&s_Ccg1="+s_Ccg1.value+"&s_Ccg2="+s_Ccg2.value+"&s_Status="+s_Status.value+"&s_Year="+s_Year.value
			}
			function AddInput(){
				window.location.href="/home/add"
			}
		</script>
	</body>
		
</html>
