package google_code_jam.`2008`.`1a`

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, ) = input_ints()
        val (M, ) = input_ints()
        val st = MutableList(N){0}
        val rows = (1..M).map{input_ints()}
        var impossible = false
        var updated = true
        while(!impossible && updated){
            updated = false
            for (row in rows) {
                var k = -1
                var ok = false
                for(i in 1 until row.size step 2){
                    val (j, v) = row[i]-1 to row[i+1]
                    if(st[j] == v){
                        ok = true
                        break
                    }
                    if(v == 1)
                        k = j
                }
                if(ok)
                    continue
                if(k == -1) {
                    impossible = true
                }else{
                    updated = true
                    st[k] = 1
                }
                break
            }
        }
        val res = if(impossible) "IMPOSSIBLE" else st.joinToString(" ")
        println("Case #${it}: $res")
    }
}
