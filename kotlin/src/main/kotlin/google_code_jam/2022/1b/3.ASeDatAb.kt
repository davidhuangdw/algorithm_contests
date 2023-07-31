package google_code_jam.`2022`.`1b`
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val M = 8
    fun getMin(s: String): String {
        var min = s
        for(i in s.indices){
            val x =  s.substring(i, s.length) + s.substring(0, i)
            if(x < min)
                min = x
        }
        return min
    }
    fun perm(k: Int): List<String> {
        val all = hashSetOf<String>()
        val path = mutableListOf<Int>()
        fun dfs(i: Int, rem: Int){
            if(rem < 0 || M-i < rem)
                return
            if(i >= M){
                all.add(getMin(path.joinToString("")))
                return
            }
            for(v in 0..1){
                path.add(v)
                dfs(i+1, rem - v)
                path.removeLast()
            }
        }
        dfs(0, k)
        return all.toList()
    }
    val P = List(9){perm(it)}
    val T = readLine()!!.toInt()
    (1..T).forEach {
        var s = P[1][0]
        var succ = false
        for(i in 0 until 300){
            println(s)
            val (res, ) = input_ints()
            if(res == 0) {
                succ = true
                break
            }
            s = P[res].random()
        }
        // todo: set2
        if(!succ) // bug: forget use if
            throw RuntimeException("failed")
    }
}
