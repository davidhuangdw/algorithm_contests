package google_code_jam.`2018`.`1c`
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (T, ) = input_ints()
    (1..T).forEach {
        val (N, L) = input_ints()
        val W = (1..N).map{split_input()[0]}

        fun set(): String{
            val have = (0 until L).map{i ->
                W.map{it[i]}.toHashSet().toList()
            }
            val all = W.toHashSet()
            var fnd = "-"
            val path = mutableListOf<Char>()
            fun dfs(i: Int): Boolean{
                if(i >= L){
                    val w = path.joinToString("")
                    if(w !in W) {
                        fnd = w
                        return true
                    }
                    return false
                }
                for(ch in have[i]){
                    path.add(ch)
                    if(dfs(i+1))
                        return true
                    path.removeLast()
                }
                return false
            }
            dfs(0)
            return fnd
        }
        val res = set()
        println("Case #${it}: $res")
    }
}
