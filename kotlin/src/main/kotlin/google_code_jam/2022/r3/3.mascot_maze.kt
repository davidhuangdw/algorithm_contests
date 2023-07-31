package google_code_jam.`2022`.r3

fun main() {
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull { if (it != "") it.toInt() else null }
    fun input_longs() = split_input().mapNotNull { if (it != "") it.toLong() else null }

    val (T) = input_ints()
    val IMP = "IMPOSSIBLE"
    val CH = "ACDEHIJKMORST"
    (1..T).forEach {
        val (N) = input_ints()
        val L = input_ints().map { listOf(it - 1) }
        val R = input_ints().map { listOf(it - 1) }
        fun set1(): String {
            val IL = MutableList(N) { mutableListOf<Int>() }
            val IR = MutableList(N) { mutableListOf<Int>() }
            for ((i, _v) in L.withIndex()) {
                for(v in _v)
                    IL[v].add(i)
            }
            for ((i, _v) in R.withIndex()) {
                for(v in _v)
                    IR[v].add(i)
            }
            val diff = MutableList(N) { hashSetOf<Int>() }
            for (i in 0 until N) {
                for((l, r) in listOf(L to R, IL to IR)){
                    val layer = List(3) { hashSetOf<Int>() }
                    layer[0].add(i)
                    for (k in 0 until 2) {
                        for (j in layer[k]) {
                            for(x in l[j])
                                layer[k + 1].add(x)
                            for(x in r[j])
                                layer[k + 1].add(x)
                        }
                    }
                    diff[i] += (layer[1] + layer[2]) as HashSet
                }
                if (i in diff[i])
                    return IMP
            }
//            println(L)
//            println(R)
//            println(IL)
//            println(IR)
//            println(diff)
            val col = MutableList(N) { -1 }
            for (i in 0 until N) {
                val exist = diff[i].map { col[it] }.toHashSet()
                for (c in 0 until 13)
                    if (c !in exist) {
                        col[i] = c
                        break
                    }
            }
            fun check(){
                for(i in 0 until N)
                    for(j in L[i]+R[i])
                        for(k in L[j] +R[j])
                            if(hashSetOf(col[i], col[j], col[k]).size != 3){
                                throw RuntimeException("failed")
                            }
            }
//            check()
            return col.map { CH[it] }.joinToString("")
        }

        val r = set1()
        println("Case #${it}: $r")
    }
}
