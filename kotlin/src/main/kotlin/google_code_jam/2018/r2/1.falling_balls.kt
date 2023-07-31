package google_code_jam.`2018`.r2
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (T, ) = input_ints()
    (1..T).forEach {
        val (C, ) = input_ints()
        val B = input_ints()
        fun set(): String {
            if(B[0] == 0 || B[C-1] == 0)
                return "IMPOSSIBLE"
            var l = 0
            var row = 0
            var sum = 0
            val seg = mutableListOf<List<Int>>()
            for((i, b) in B.withIndex()){
                if(b > 0){
                    sum += b
                    val r = sum-1
                    seg.add(listOf(i, l, r))
                    row = maxOf(row, i-l, r-i)
                    l = sum
                }
            }
//            println(seg)
            row ++
            val map = List(row){MutableList(C){'.'} }
            for((md, l, r) in seg){
                for((i, j) in (l..md-1).withIndex()){
                    map[i][j] = '\\'
                }
                for((i, j) in (r downTo md+1).withIndex()){
                    map[i][j] = '/'
                }
            }
            return buildString {
                appendLine(row)
                for(r in map)
                    appendLine(r.joinToString(""))
            }
        }
        val r = set()
        println("Case #${it}: $r")
    }
}
