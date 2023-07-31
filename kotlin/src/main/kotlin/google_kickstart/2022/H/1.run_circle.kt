package google_kickstart.`2022`.H
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (T, ) = input_ints()
    (1..T).forEach {
        val (L, N) = input_ints()
        var cnt = 0L
        var dir = "C"
        var off = 0L
        for(i in 0 until N){
            val (l, d) = split_input()
            var len = l.toLong()
            if(off == 0L || dir == d){
                cnt += (len + off) / L
                off = (len+off) % L
                dir = d
            }else if(len < off){
                off -= len
            }else{
                len -= off
                cnt += len / L
                off = len % L
                dir = d
            }
        }
        println("Case #${it}: $cnt")
    }
}
