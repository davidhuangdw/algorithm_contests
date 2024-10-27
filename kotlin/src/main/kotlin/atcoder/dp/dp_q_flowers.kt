package atcoder.dp
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (N, ) = input_ints()
    val H = input_ints()
    val A = input_longs()
    val pre_max = LongArray(N+1)
    H.zip(A).forEach{(h, a)->
        var t = h;
        var mx = 0L
        while (t > 0){
            mx = maxOf(mx, pre_max[t])
//            println("--- ${i} ${h}: ${t} ${mx}")
            t = (t and (t+1)) - 1
        }
        val cur = mx+a
        pre_max[h] = cur

        t = h
        while(t <= N){
            pre_max[t] = maxOf(pre_max[t], cur)
            t = t or (t+1)
        }
//        println("=== ${i} ${h}: ${pre_max.toList()}")
    }
//    println(pre_max.toList())
    println(pre_max.maxOrNull()!!)
}

