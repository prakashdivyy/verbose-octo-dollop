$("#search").submit(function (e) {
    e.preventDefault();
    var form = $("#search");
    var qs = form.find("input[name='qs']").val();
    var req = $.post("/", { qs: qs });
    req.done(function (data) {
        mhs = data.data;
        banyak = mhs.length;
        if (banyak > 0) {
            $("#pala").empty();
            $("#badan").empty();
            $("#count").text("Ditemukan " + banyak + " data mahasiswa yang cocok.");
            var header = "<tr><td>NPM</td><td>Nama</td><td>Email</td><td>Fakultas</td></tr>"
            $("#hasil thead").append(header);
            for (var ii = 0; ii < banyak; ii++) {
                var nama = mhs[ii].nama;
                var npm = mhs[ii].npm;
                var email = mhs[ii].email;
                var fakultas = mhs[ii].fakultas;
                var content = "<tr><td><a href='/foto/" + npm + "'>" + npm + "</a></td><td>" + nama + "</td><td>" + email + "</td><td>" + fakultas + "</td></tr>";
                $("#hasil tbody").append(content);
            }
        }
        else {
            $("#count").text("Tidak ada data mahasiswa yang cocok.");
            $("#pala").empty();
            $("#badan").empty();
        }
    });
});