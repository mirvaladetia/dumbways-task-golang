let dataBlog=[];
function addBlog(event) {
    event.preventDefault();
    let judul=document.getElementById("input-judul-cerita").value
    let waktumulai=document.getElementById("input-start-date").value;
    let waktuberakhir=document.getElementById("input-end-date").value;
    let konten=document.getElementById("input-konten").value
    let role=document.getElementsByName("role")
    function mydate(date)
    {
      var
        month = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'],
        days  = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday']
      ;
      return days[date.getDay()]+' '+month[date.getMonth()]+' '+date.getDate()+' '+date.getFullYear()
    }
    var result = []
    for (var i=0; i < role.length; i++) {
        if (role[i].checked) {
            result.push(`<img src=${role[i].value} alt="gambar-node" />`)
        } else {
            result.push("")
        }
        
    }
    let foto=document.getElementById("unggah-foto-blog").files

    if (foto.length == 0) {
        return alert("Harap unggah foto terlebih dahulu")
    }
    foto = URL.createObjectURL(foto[0])
    
let durasi=getDifferenceTimeInput(waktumulai, waktuberakhir)
let blog = {
    judul, konten, durasi, result, foto, postAt : mydate(new Date)
}
dataBlog.push(blog)
renderBlog()
}
function getDifferenceTimeInput(waktumulai, waktuberakhir) {
    let startToMilliseconds = new Date(waktumulai).getTime();
    console.log(startToMilliseconds);
    let endToMilliseconds = new Date(waktuberakhir).getTime();
    let selisih = endToMilliseconds - startToMilliseconds;
    let durasiHari = Math.floor(selisih / (1000 * 60 * 60 * 24));
    let durasiMinggu = Math.floor(selisih / (1000 * 60 * 60 * 24 * 7));
    let durasiBulan = Math.floor(selisih / (1000 * 60 * 60 * 24 * 7 * 4));
    let durasiTahun = Math.floor(selisih / (1000 * 60 * 60 * 24 * 7 * 4 * 12));

    if (durasiTahun > 0) {
      return `${durasiTahun} Tahun / ${durasiBulan} Bulan / ${durasiMinggu} Minggu / ${durasiHari} Hari`;
    } else if (durasiBulan > 0) {
      return `${durasiBulan} Bulan / ${durasiMinggu} Minggu / ${durasiHari} Hari`;
    } if (durasiMinggu > 0) {
      return `${durasiMinggu} Minggu / ${durasiHari} Hari`;
    } if (durasiHari > 0) {
      return `${durasiHari} Hari`
    }
  }
function renderBlog (){
    document.getElementById("contents").innerHTML =""
    for (let index = 0; index < dataBlog.length; index++) {
        document.getElementById("contents").innerHTML += ` 
        <div class="blog-list-container">
        <div class="blog-list-image">
            <img src="${dataBlog[index].foto}" alt="blog-gambar"/>
        </div>
        <div class="blog-content">
            <div class="button-group">
                <button class="button-unggah">Unggah Post</button>
                <button class="button-hapus">Hapus Post</button>
            </div>
            <h1>
                <a href="blog_list.html" target="blank">
                ${dataBlog[index].judul}</a></h1>
            <div style="display: flex; align-items: center; justify-content: right">
            <div style="margin-right: 10px;">
            ${dataBlog[index].result[0]}
            ${dataBlog[index].result[1]}
            ${dataBlog[index].result[2]}
            ${dataBlog[index].result[3]}
            </div>
            <div>
            <p>durasi : ${dataBlog[index].durasi}</p>
            </div>
            </div> 
            <p class="post-uploader">${dataBlog[index].postAt} | By : Mirval Adetia</p>
            <div>
            <p>
            ${dataBlog[index].konten}
            </p>
        </div>
    </div> 
        `
    }
}
// function getFullTime(time) {
//     // console.log("get full time");
//     // let time = new Date();
//     // console.log(time);
  
//     let monthName = [
//       "Jan",
//       "Feb",
//       "Mar",
//       "Apr",
//       "May",
//       "Jun",
//       "Jul",
//       "Aug",
//       "Sep",
//       "Oct",
//       "Nov",
//       "Dec",
//     ];
//     // console.log(monthName[8]);
  
//     let date = time.getDate();
//     // console.log(date);
  
//     let monthIndex = time.getMonth();
//     // console.log(monthIndex);
  
//     let year = time.getFullYear();
//     // console.log(year);
  
//     let hours = time.getHours();
//     let minutes = time.getMinutes();
//     // console.log(minutes);
  
//     if (hours <= 9) {
//       hours = "0" + hours;
//     } else if (minutes <= 9) {
//       minutes = "0" + minutes;
//     }
  
//     return `${date} ${monthName[monthIndex]} ${year} ${hours}:${minutes} WIB`;
//   }
  
//   function getDistanceTime(time) {
//     let timeNow = new Date();
//     let timePost = time;
  
//     // waktu sekarang - waktu post
//     let distance = timeNow - timePost; // hasilnya milidetik
//     console.log(distance);
  
//     let milisecond = 1000; // milisecond
//     let secondInHours = 3600; // 1 jam 3600 detik
//     let hoursInDays = 24; // 1 hari 24 jam
  
//     let distanceDay = Math.floor(
//       distance / (milisecond * secondInHours * hoursInDays)
//     ); // 1/86400000
//     let distanceHours = Math.floor(distance / (milisecond * 60 * 60)); // 1/3600000
//     let distanceMinutes = Math.floor(distance / (milisecond * 60)); // 1/60000
//     let distanceSeconds = Math.floor(distance / milisecond); // 1/1000
  
//     if (distanceDay > 0) {
//       return `${distanceDay} Day Ago`;
//     } else if (distanceHours > 0) {
//       return `${distanceHours} Hours Ago`;
//     } else if (distanceMinutes > 0) {
//       return `${distanceMinutes} Minutes Ago`;
//     } else {
//       return `${distanceSeconds} Seconds Ago`;
//     }
//   }
  
//   setInterval(function () {
//     renderBlog();
//   }, 3000);