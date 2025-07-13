package atcoder.a6

fun main(){
    fun input_ints() = readLine()!!.split(" ").map { it.toInt() }

    fun top_thr_bits(v: Long): Long{
        val bits = mutableListOf<Long>()
        var r = v
        while(r > 0){
            val b = r and -r
            bits.add(b)
            r -= b
        }
        if(bits.size < 3){
            return -1L
        }
        return bits.reversed().take(3).reduce(Long::plus)
    }

    val (T, ) = input_ints()
    repeat(T){
        val S = readLine()!!
        var n = S.toLong()
        if(n < 7){
            println(-1)
        }else{
            var found = -1L
            for(i in 0 until 4){
                val t = top_thr_bits(n)
                if(t > 0){
                    found = t
                    break
                }
                n--
            }
            println(found)
        }
    }
}