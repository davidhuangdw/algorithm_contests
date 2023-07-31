package google_code_jam.`2019`.`1b`
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    fun binary(x: Long){
        val s = (63 downTo 0).map{if(x and (1L shl it) > 0) 1 else 0}.joinToString("")
//        println(s)
    }
    fun pick(x: Long, k: Int) = ((x shr k) % 128).toInt()
    val (T, W) = input_ints()
    (1..T).forEach {
        val R = MutableList(7){0}
        var d = 4 * 56
        println(d)
        var (v, ) = input_longs()
        binary(v)
        R[4] = pick(v, d/4)
        R[5] = pick(v, d/5)
        R[6] = pick(v, d/6)

        d = 56
        println(d)
        v = input_longs()[0]
        for(i in 4..6) // bug: should minus first, because R[5], R[6] added could increase R[4] value!!
            v -= (R[i].toLong()) shl d/i
        binary(v)
        R[1] = pick(v, d)
        R[2] = pick(v, d/2)
        R[3] = pick(v, d/3) // bug: shr must in bracket!! priority < "+, -"
//        println(R)
        val res = R.subList(1, 7).joinToString(" ")
        println(res)
        val (succ, ) = split_input()
        if(succ != "1")
            throw RuntimeException("failed")
//        println("Case #${it}: $res")
    }
}
