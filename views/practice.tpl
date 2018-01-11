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
							<h4 class="panel-title">查询实践环节信息</h4>
						</div>
					    <div class="panel-body">
						<form class="form-inline" role="form" id="searchPractive">
					        <div class="form-group">				
								<label>开课单位</label>			
								<input class="form-control" name="s_Punit" id="s_Punit">
								<label>环节名称</label>			
								<input class="form-control" name="s_Pname" id="s_Pname">
								<button type="button" class="btn btn-default" onclick="return QueryInput()">检索</button>				
							</div>
						</form>
						<div class="col-sm-1 pull-right">					
							<button type="button" class="btn btn-primary" onclick="return AddInput()">新增</button>																	 	
					    </div>
						</div>
					</div>
					<div class="table-responsive">
						<table class="table table-bordered">
							<caption><h4 class="panel-title">实践环节信息</h4></caption>
							<thead>
								<tr>
									<th>操作</th>
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
									<th><a href="/practice/edit?pid={{.Pid}}">编辑</a></th>
									<th>{{.Pid}}</th>
									<th>{{.Punit}}</th>
									<th>{{.Pname}}</th>
									<th>{{.Pcg1}}</th>
									<th>{{.Pname_en}}</th>
									<th>{{.Status}}</th>
									<th>{{.Credit}}</th>
									<th>{{.Tclass}}</th>
									<th>{{.Nw}}</th>>
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
				var s_Punit=document.getElementById("s_Punit")
				var s_Pname=document.getElementById("s_Pname")
				window.location.href="/practice/search?s_Punit="+s_Punit.value+"&s_Pname="+s_Pname.value
			}
			function AddInput(){
				window.location.href="/practice/add"
			}
		</script>
	</body>
		
</html>
