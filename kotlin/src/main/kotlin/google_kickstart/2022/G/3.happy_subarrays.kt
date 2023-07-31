package google_kickstart.`2022`.G

import java.util.TreeSet

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (T, ) = input_ints()
    (1..T).forEach {
        val (N, ) = input_ints()
        val A = input_longs()
        fun set1(): Long{
            var sum = 0L
            for(i in 0 until N){
                var acc = 0L
                for (j in i until N){
                    acc += A[j]
                    if(acc < 0)
                        break
                    sum += acc
                }
            }
            return sum
        }

        // good problem: index
        fun set2(): Long{
            val B = (0 until N).map{A[it] * (N-it)}
            val prea = A.scan(0, Long::plus)
            val preb = B.scan(0, Long::plus)

            val st = mutableListOf(N)
            var res = 0L
            for(i in N-1 downTo 0){
                // pop
                while(st.isNotEmpty() && prea[i] <= prea[st.last()])
                    st.removeLast()
                val j = if(st.isEmpty()) N else st.last()-1
                res += preb[j] - preb[i] - (N-j) * (prea[j] - prea[i])
                // insert
                st.add(i)
            }
            return res
        }
        println("Case #${it}: ${set2()}")
    }
}
