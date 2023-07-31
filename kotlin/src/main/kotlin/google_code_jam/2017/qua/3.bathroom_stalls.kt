package google_code_jam.`2017`.qua

import java.util.*

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (T, ) = input_ints()
    (1..T).forEach {
        val (N, K) = input_longs()
        fun maxMin(len: Long): String{
            return "${len/2} ${(len-1)/2}"
        }
        // bug: every layer, not just two values!!!
        fun set2(): String {
            val que = PriorityQueue<Long>(compareByDescending{it})
            que.add(N)
            for(_i in 0 until K.toInt() -1){
                val s = que.poll() - 1
                que.add((s+1)/2)
                que.add(s/2)
            }
            return maxMin(que.poll())
        }
        fun set3(): String {
            val cnt = hashMapOf<Long, Long>()  // bug: value is long!!!
            var r = K
            val que = PriorityQueue<Long>(compareByDescending{it})
            que.add(N)
            cnt[N] = 1
            while(true){
                val s = que.poll()
                val c = cnt[s]!!
                if(r <= c)
                    return maxMin(s)
                r -= c
                for(h in listOf(s/2, (s-1)/2)){
                    if(h !in cnt){
                        que.add(h)
                        cnt[h] = 0
                    }
                    cnt[h] = cnt[h]!! + c
                }
            }
        }
        val res = set3()
        println("Case #${it}: $res")
    }
}
