let dataBlog = []

function addBlog(event) {
    event.preventDefault()

    let project = document.getElementById("input-project").value
    let starDate = document.getElementById("input-star-date").value
    let endDate = document.getElementById("input-end-date").value
    let content = document.getElementById("input-content").value
    let cekbox1 = document.getElementById("checkbox1").value
    let cekbox2 = document.getElementById("checkbox2").value
    let cekbox3 = document.getElementById("checkbox3").value
    let cekbox4 = document.getElementById("checkbox4").value
    let image  = document.getElementById("input-blog-image").files[0]

    // buat url gambar nantinya tampil

    image = URL.createObjectURL(image)
    //console.log(image)

    let blog = {
        project,
        starDate: new Date(),
        endDate,
        content,
        cekbox1,
        cekbox2,
        cekbox3,
        cekbox4,
        image,
        author: " Yoga "
    }

    dataBlog.push(blog)
    //console.log(dataBlog)

    renderBlog()
}

function renderBlog() {
    document.getElementById("contents").innerHTML = ''

    for (let index = 0; index < dataBlog.length; index++) {
        console.log("test",dataBlog[index])

        document.getElementById("contents").innerHTML += `
        <div class="cards">
        <div class="card-img">
            <img src="${dataBlog[index].image}">
            
        </div>
        <div class="card-body">
            <a href="blog-detail.html" target="_blank">
            ${dataBlog[index].project}</a>

            <div style="font-size: 12px; color:grey">
            <p>Durasi </p>
             Post in:${getFullTime(dataBlog[index].starDate)} | Last Post: ${getDistanceTime(dataBlog[index].starDate)}
            </div>
            
        
            
            <p>${dataBlog[index].content}</p>

                <p style=" margin-left:10px;  text-align:left;padding:15px;">

                    ${dataBlog[index].cekbox1} ${dataBlog[index].cekbox2} ${dataBlog[index].cekbox3} ${dataBlog[index].cekbox4}
                        
                </p>
        </div>
        <div class="card-footer">
            <button>Edit</button>
            <button>Delete</button>
        </div>
        <div>
             <p style="text-align: center; font-size: 15px; color: black">Copyright 2022. :${dataBlog[index].author}</p>
        </div>
    </div>

        `
    }
}

    //input tanggal awal
function getFullTime(time) {
    // time = new Date()
    // console.log(time)
     let x = ("car","moto")
    let monthName = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']
    // console.log(monthName[9])

    // 14
    let date = time.getDate()
    console.log(date)

    // 9
    let monthIndex = time.getMonth()
    console.log(monthIndex)

    // 2022
    let year = time.getFullYear()
    console.log(year)

    let hours = time.getHours()
    let minutes = time.getMinutes()

    console.log(hours)


    if (hours <= 9) {
        hours = "0" + hours
    } 
    
    if (minutes <= 9) {
        minutes = "0" + minutes
    }

    // 14 Oct 2022 09:07 WIB
    return `${date} ${monthName[monthIndex]} ${year} ${hours}:${minutes} WIB`
}

function getDistanceTime(time) {
    let timeNow = new Date()
    let timePost = time

    let distance = timeNow - timePost //milisecond
    console.log(distance)

    let milisecond = 1000 // milisecond
    let secondInHours = 3600 // 1 jam = 3600 detik
    let hoursInDay = 24 // 1 hari = 24 jam

    let distanceDay = Math.floor(distance / (milisecond * secondInHours * hoursInDay))
    let distanceHours = Math.floor(distance / (milisecond * 60 * 60))
    let distanceMinutes = Math.floor(distance / (milisecond * 60))
    let distanceSecond = Math.floor(distance / milisecond)

    if (distanceDay > 0) {
        return `${distanceDay} day(s) ago`
    } else if (distanceHours > 0) {
        return `${distanceHours} hour(s) ago`
    } else if (distanceMinutes > 0) {
        return `${distanceMinutes} minute(s) ago`
    } else {
        return `${distanceSecond} second(s) ago`
    }
}


setInterval(function() {
    renderBlog()
}, 3000)
