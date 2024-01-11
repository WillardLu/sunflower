// @copyright Copyright 2024 Willard Lu
// @email willard.lu@outlook.com
// @language JavaScript
// @author 陆巍
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.
// 打开对话框
function openDialog() {
  document.querySelector("#myDialog").show();
}

// 关闭对话框
function closeDialog() {
  document.querySelector("#myDialog").close();
}

// 提交数据
function submitData() {
  const xhr = new XMLHttpRequest();
  xhr.open("POST", "/submit-data");
  xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  const body = JSON.stringify({
    "task_number": document.getElementById("task-number").value,
    "task_name": document.getElementById("task-name").value
  })
  xhr.onload = () => {
    if (xhr.readyState == 4 && xhr.status == 201) {
      //var data = JSON.parse(xhr.responseText); // 注意json的键名中不要包含减号“-”
      var data = xhr.responseText
      //alert(data.task_number + ": " + data.task_name);
      alert(data);
      console.log(data);
    } else {
      console.log("Error: ${xhr.status}");
    }
  }
  xhr.send(body);
  document.getElementById("task-number").value = "";
  document.getElementById("task-name").value = "";
  closeDialog();
}
