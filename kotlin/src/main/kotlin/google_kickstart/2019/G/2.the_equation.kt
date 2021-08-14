package google_kickstart.`2019`.G

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val B = 50
        val (N, M) = input_longs()
        val bits_count = List(B){ MutableList(2){0}}
        for(a in input_longs()) {
            for (i in 0 until B) {
                val b = if (a and (1L shl i) > 0) 1 else 0
                bits_count[i][b] += 1
            }
        }

        fun get_max(): Long{
            var base = 0L
            val min_count = (0 until B).map{ i ->
                val m = bits_count[i].minOrNull()!!
                base += (1L shl i) * m
                m
            }
            if(base > M) return -1

            var x = 0L
            for(i in B-1 downTo 0){
                val p = 1L shl i
                val b = if((bits_count[i][0] - min_count[i]) * p + base <= M) 1 else 0
                base += (bits_count[i][b xor 1] - min_count[i]) * p
                if(b == 1)
                    x = x or p
            }
            return x
        }

        println("Case #${it}: ${get_max()}")
    }
}
