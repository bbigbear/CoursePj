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
								<label>承担单位</label>			
								<input class="form-control" name="s_Cunit">
								<label>课程名称</label>			
								<input class="form-control" name="s_Cname">
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
							<caption><h4 class="panel-title">理论课程信息</h4></caption>
							<thead>
								<tr>
									<th>操作</th>
									<th>课程代码</th>
									<th>承担单位</th>
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
									<th><a href="/home/edit?cid={{.Cid}}">编辑</a></th>
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
				$.ajax({
					type:"POST",
					url:"{{urlfor "HomeController.TheoryCourseSearch"}}",
					data:$("#searchCourse").serialize(),
					async:false,
					error:function(request){
						alert("post error")
						
					},
					success:function(data){
						if(data.status==0){
							alert("查询成功")
						}else{
							alert("查询失败")
						}
						
					}
					
				});
				return true	
			}
			function AddInput(){
				window.location.href="/home/add"
			}
		</script>
	</body>
		
</html>
