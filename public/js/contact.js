function ngumpulindata() {
    let name=document.getElementById("input-name").value
    let email=document.getElementById("input-email").value
    let phone=document.getElementById("input-phone").value
    let subject=document.getElementById("input-subject").value
    let textarea=document.getElementById("input-message").value
    if (name=="") {return alert("Nama harus diisi")}
    if (email=="") {return alert("Email harus diisi")}
    if (phone=="") {return alert("Nomor Handphone harus diisi")}
    if (subject=="") {return alert("Subject harus dipilih")}
    if (textarea=="") {return alert ("Pesan Harus diisi")}

let kirimemail="mirvaladetia@gmail.com"
let pengirim=document.createElement('a')

pengirim.href=`mailto:${kirimemail}?subject=${subject}&body=Halo, Perkenalkan 
nama saya ${name}, ${textarea}. Jika berkenan mohon menghubungi saya di ${phone}`
pengirim.click()
let pengirimemail ={
    name, email, phone, subject, textarea
}
}
