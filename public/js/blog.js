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

// function confirmation(){
//   var result = confirm("Are you sure to delete?");
//   if(result){
//     console.log("Deleted")
//   }
// }