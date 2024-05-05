 function AESencryption() {
                var content = document.getElementById("contentAES").value;

                fetch('http://localhost:8080/AES/encryption', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ content: content })
                })
                    .then(response => response.json())
                    .then(data => {
                        document.getElementById("resultsAES").innerText = data.encrypted_message;
                        document.getElementById("keyAES").value = data.key;
                    })
                    .catch(error => console.error('Error:', error));
            }
            function AESdecryption() {
                var content = document.getElementById("contentAES").value;
                var key = document.getElementById("keyAES").value;

                fetch('http://localhost:8080/AES/decryption', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ content: content, key: key })
                })
                    .then(response => response.json())
                    .then(data => {
                        document.getElementById("resultsAES").innerText = data.decrypted_message;
                    })
                    .catch(error => console.error('Error:', error));

                    console.log(key)
            }

            function RSAencryption() {
                    var content = document.getElementById("contentRSA").value;

                    fetch('http://localhost:8080/RSA/encryption', {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json'
                        },
                        body: JSON.stringify({ content: content })
                    })
                        .then(response => response.json())
                        .then(data => {
                            document.getElementById("resultsRSA").innerText = data.encrypted_message;
                            document.getElementById("publicKeyRSA").innerText = data.publicKey;
                            document.getElementById("privateKeyRSA").innerText = data.privateKey;
                        })
                        .catch(error => console.error('Error:', error));
                }

                function RSAdecryption() {
                    var content = document.getElementById("contentRSA").value;
                    var privateKey = document.getElementById("privateKeyRSA").innerText;

                    fetch('http://localhost:8080/RSA/decryption', {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json'
                        },
                        body: JSON.stringify({ content: content, private_key: privateKey })
                    })
                        .then(response => response.json())
                        .then(data => {
                            document.getElementById("resultsRSA").innerText = data.decrypted_message;
                        })
                        .catch(error => console.error('Error:', error));
                }

                function deleteRecord(id) {
                        fetch('http://localhost:8080/delete', {
                            method: 'POST',
                            headers: {
                                'Content-Type': 'application/json'
                            },
                            body: JSON.stringify({ id: id })
                        })
                            .then(response => {
                                if (!response.ok) {
                                    throw new Error('Failed to delete record');
                                }
                                return response.json();
                            })
                            .then(data => {
                                alert(data.message); // Hiển thị thông báo khi xóa thành công
                                // Gọi lại hàm readAllData để cập nhật lại dữ liệu sau khi xóa
                                readAllData();
                            })
                            .catch(error => {
                                console.error('Error:', error);
                                // Xử lý lỗi nếu cần
                            });
                    }

                function readAllData() {
                    fetch('http://localhost:8080/readalldata')
                        .then(response => {
                            if (!response.ok) {
                                throw new Error('Failed to fetch data');
                            }
                            return response.json();
                        })
                        .then(data => {
                            renderData(data);
                        })
                        .catch(error => {
                            console.error('Error:', error);
                            // Xử lý lỗi, nếu cần
                        });
                }

                function renderData(data) {
                    const tbody = document.getElementById("dataRows");
                    tbody.innerHTML = ""; // Xóa các hàng dữ liệu cũ trước khi thêm dữ liệu mới

                    data.forEach(item => {
                        const newRow = document.createElement("tr");

                        newRow.innerHTML = `
                            <td>${item.id}</td>
                            <td>${item.type}</td>
                            <td>${item.key}</td>
                            <td>${item.public_key}</td>
                            <td>${item.private_key}</td>
                            <td>${item.content}</td>
                            <td>${item.encrypted_content} <br><button class="btn btn-success" onclick="copyEn('${item.encrypted_content}')">Copy</button></td>
                            <td>${item.created_at}</td>
                            <td><button class="btn btn-danger" onclick="deleteRecord('${item.id}')">Delete</button></td>

                        `;

                        tbody.appendChild(newRow);
                    });
                }
                function resetFields() {
                    // Reset input và span của RSA
                    document.getElementById("contentRSA").value = "";
                    document.getElementById("publicKeyRSA").innerText = "";
                    document.getElementById("privateKeyRSA").innerText = "";
                    document.getElementById("resultsRSA").innerText = "";
                }
                function resetField() {
                    // Reset input và span của AES
                    document.getElementById("contentAES").value = "";
                    document.getElementById("keyAES").innerText = "";
                    document.getElementById("resultsAES").innerText = "";
                }

                function copyContent() {
                    // Get the content of the span element
                    var content = document.getElementById("resultsAES").innerText;

                    // Create a temporary input element
                    var tempInput = document.createElement("input");

                    // Set the value of the input to the content we want to copy
                    tempInput.setAttribute("value", content);

                    // Append the input element to the DOM
                    document.body.appendChild(tempInput);

                    // Select the content of the input element
                    tempInput.select();

                    // Copy the selected content
                    document.execCommand("copy");

                    // Remove the temporary input element
                    document.body.removeChild(tempInput);

                    alert("Copy success!!!");
                }
                function copyContents() {
                    // Get the content of the span element
                    var content = document.getElementById("resultsRSA").innerText;

                    // Create a temporary input element
                    var tempInput = document.createElement("input");

                    // Set the value of the input to the content we want to copy
                    tempInput.setAttribute("value", content);

                    // Append the input element to the DOM
                    document.body.appendChild(tempInput);

                    // Select the content of the input element
                    tempInput.select();

                    // Copy the selected content
                    document.execCommand("copy");

                    // Remove the temporary input element
                    document.body.removeChild(tempInput);

                    alert("Copy success!!!");
                }
               function copyEn(encryptedContent) {
                    // Tạo một thẻ <textarea> tạm thời để chứa nội dung
                    var tempTextArea = document.createElement("textarea");
                    tempTextArea.value = encryptedContent;

                    // Thêm thẻ vào trang web
                    document.body.appendChild(tempTextArea);

                    // Chọn toàn bộ nội dung trong thẻ <textarea>
                    tempTextArea.select();
                    tempTextArea.setSelectionRange(0, 999999); // Đối với một số trình duyệt

                    // Sao chép nội dung vào clipboard
                    document.execCommand("copy");

                    // Xóa thẻ <textarea> tạm thời
                    document.body.removeChild(tempTextArea);

                    // Thông báo sau khi sao chép thành công (tùy chọn)
                    alert("Encrypted content copied to clipboard!");
                }

