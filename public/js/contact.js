//Fungsi memanggil data dari html berdasarkan id (ex: input-name)
function showData (){
    let showName = document.getElementById('input-name').value;
    let showEmail = document.getElementById('input-email').value;
    let showPhone = document.getElementById('input-phone').value;
    let showSubject = document.getElementById('input-subject').value;
    let showMessage = document.getElementById('input-message').value;

    console.log(showName)
    console.log(showEmail)
    console.log(showPhone)
    console.log(showSubject)
    console.log(showMessage)

   
   
   //pengkondisian form

    if (showName == ''){
        return alert ('Nama harus di isi')
    }

    else if (showEmail == ''){
        return alert ('email harus di isi')
    }

    else if (showPhone == ''){
        return alert ('Phone harus di isi')
    }

     else if (showSubject == ''){
        return alert ('Subject harus di isi')
    }

     else if (showMessage == ''){
        return alert ('Message harus di isi')
    }
    
    let emailReceiver = 'yogaputera0@gmail.com'

    let a = document.createElement('a');

    a.href = `mailto:${emailReceiver}?subject:${showSubject}&body= Hello, my name is ${showName}, ${showMessage}`
    a.click()
}   