package google_code_jam.`2018`.`1b`

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (T) = input_ints()
    (1..T).forEach {
        val (S) = input_ints()
        val A = (1..S).map {
            val (d, a, b) = input_ints()
            listOf(d + a, d - b)
        }.toMutableList()

        fun set1(): String {
            var max = 0
            var st = hashSetOf<Int>()  // bug: different(m,n) could have same continuous_set
            for (m in A.map { it[0] }.toHashSet())
                for (n in A.map { it[1] }.toHashSet()) {
                    var l = S
                    for ((i, x) in A.withIndex()) {
                        val (a, b) = x
                        if (a == m || b == n) {
                            l = minOf(l, i)
                            val d = i - l + 1
                            if (d > max) {
                                max = d
                                st = hashSetOf(l)
                            } else if (d == max) {
                                st.add(l)
                            }
                        } else {
                            l = S
                        }
                    }
                }
            return "$max ${st.size}"
        }

        fun set2(): String {
            // bug: only conside mmmnnn case, not aware the case of mnmnmn
            var max = 0
            var cnt = 0
            val mr = MutableList(S+1){S}
            val nr = MutableList(S+1){S}
            val mnr = MutableList(S+1){S}
            val nmr = MutableList(S+1){S}
            val vm = MutableList(S+1){S}  // redundant for easy code
            val vn = MutableList(S+1){S}
            A.add(listOf(0, 0))
//            println(A)
            val M = A.map{it[0]}
            val N = A.map{it[1]}
            for(i in S-1 downTo 0){
                mr[i] = if(M[i] == M[i+1]) mr[i+1] else i+1
                vn[i] = N[mr[i]]
                nr[i] = if(N[i] == N[i+1]) nr[i+1] else i+1
                vm[i] = M[nr[i]]

                mnr[i] = if(vm[mr[i]] == M[i]) nmr[mr[i]] else nr[mr[i]]
                nmr[i] = if(vn[nr[i]] == N[i]) mnr[nr[i]] else mr[nr[i]]
                val x = maxOf(mnr[i], nmr[i]) - i
//                println("$i: $x")
                if(x > max){
                    max = x
                    cnt = 1
                }else if (x == max){
                    cnt++
                }
            }

            return "$max $cnt"
        }

        val res = set2()
        println("Case #${it}: $res")
    }
}
